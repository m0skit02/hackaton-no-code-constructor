package main

import (
	"context"
	"fmt"
	hackaton_no_code_constructor "hackaton-no-code-constructor"
	"hackaton-no-code-constructor/pkg/handler"
	models "hackaton-no-code-constructor/pkg/model"
	"hackaton-no-code-constructor/pkg/repository"
	"hackaton-no-code-constructor/pkg/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gorm.io/gorm"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	// Попробуем загрузить .env из текущей или родительской папки
	if err := loadEnv(); err != nil {
		logrus.Warn("⚠️  .env file not found, using Docker environment variables")
	} else {
		logrus.Info("✅ .env file loaded successfully")
	}

	db := connectToPostgres()
	runMigrations(db)

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	// Контекст для graceful shutdown
	ctx, shutdown := waitForShutdown()
	defer shutdown()

	srv := new(hackaton_no_code_constructor.Server)
	go func() {
		address := buildAppAddress()
		logrus.Infof("🚀 Starting server on %s", address)

		if err := srv.Run(address, handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("❌ Error running HTTP server: %s", err)
		}
	}()

	logrus.Info("✅ Server started successfully")

	// Ждём сигнала через контекст
	<-ctx.Done()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(shutdownCtx); err != nil {
		logrus.Errorf("error during server shutdown: %s", err)
	}

	closeDB(db)

	logrus.Info("🟢 Application shutdown complete")
}

func loadEnv() error {
	if err := godotenv.Load("./.env"); err == nil {
		return nil
	}
	if err := godotenv.Load("../.env"); err == nil {
		return nil
	}
	return fmt.Errorf(".env not found")
}

func connectToPostgres() *gorm.DB {
	requiredEnv := []string{
		"POSTGRES_HOST", "POSTGRES_PORT", "POSTGRES_USER",
		"POSTGRES_DATABASE", "POSTGRES_PASSWORD", "POSTGRES_SSL_MODE",
	}

	for _, key := range requiredEnv {
		if os.Getenv(key) == "" {
			logrus.Fatalf("environment variable %s is required but not set", key)
		}
	}

	cfg := repository.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		DBName:   os.Getenv("POSTGRES_DATABASE"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
	}

	db, err := repository.NewPostgresDB(cfg)
	if err != nil {
		logrus.Fatalf("failed to initialize Postgres: %v", err)
	}

	logrus.Info("✅ Connected to Postgres")
	return db
}

func waitForShutdown() (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-quit
		logrus.Infof("🛑 Received signal %s, shutting down...", sig)
		cancel()
	}()

	return ctx, cancel
}

func closeDB(db *gorm.DB) {
	if db == nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		logrus.Errorf("failed to get sql.DB from gorm: %s", err)
		return
	}
	if err := sqlDB.Close(); err != nil {
		logrus.Errorf("error closing DB connection: %s", err)
	} else {
		logrus.Info("🔒 PostgreSQL connection closed")
	}
}

func buildAppAddress() string {
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}
	// Всегда слушаем на всех интерфейсах
	return "0.0.0.0:" + port
}

func runMigrations(db *gorm.DB) {
	logrus.Info("Running auto-migrations...")

	if err := db.AutoMigrate(
		&models.BlockType{},
		&models.Tag{},
	); err != nil {
		logrus.Fatalf("migration failed: %v", err)
	}

	logrus.Info("Auto-migrations completed successfully")
}
