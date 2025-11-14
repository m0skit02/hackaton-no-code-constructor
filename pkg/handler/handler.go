package handler

import (
	"hackaton-no-code-constructor/pkg/middleware"
	"hackaton-no-code-constructor/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler { return &Handler{services: services} }

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	auth := router.Group("/api/auth")
	{
		auth.POST("/", h.Login)
	}

	api := router.Group("/api")
	api.Use(middleware.JWTAuthMiddleware())
	{
		blockTypes := api.Group("/blockTypes")
		{
			blockTypes.POST("/", h.createBlockType)      // Создать блок
			blockTypes.GET("/", h.getAllBlockTypes)      // Получить список блоков
			blockTypes.GET("/:id", h.getBlockTypeById)   // Получить блок по ID
			blockTypes.PUT("/:id", h.updateBlockType)    // Обновить блок по ID
			blockTypes.DELETE("/:id", h.deleteBlockType) // Удалить блок по ID
		}

		tags := api.Group("/tags")
		{
			tags.POST("/", h.createTag)      // Создать Тег
			tags.GET("/", h.getAllTags)      // Получить список тегов
			tags.PUT("/:id", h.updateTag)    // Обновить тег по ID
			tags.DELETE("/:id", h.deleteTag) // Удалить тег по ID
			tags.GET("/:id", h.getTagById)   // Получить тег по ID
		}
	}
	return router
}
