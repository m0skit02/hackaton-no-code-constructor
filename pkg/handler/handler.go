package handler

import (
	"aeza-checker/pkg/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler { return &Handler{services: services} }

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	api := router.Group("/api")
	{
		agents := api.Group("/agents")
		{
			agents.POST("/", h.createAgent)                      // создать агента
			agents.GET("/", h.getAllAgents)                      // список агентов
			agents.GET("/:id", h.getAgentByID)                   // получить агента по ID
			agents.DELETE("/:id", h.deleteAgent)                 // удалить агента
			agents.PUT("/:id/heartbeat", h.updateAgentHeartbeat) // обновить heartbeat
			agents.PUT("/:id/status", h.setAgentStatus)          // изменить статус агента
		}

		tasks := api.Group("/tasks")
		{
			tasks.POST("/", h.createTask)
			tasks.GET("/", h.getAllTask)
			tasks.DELETE("/:id", h.deleteTask)
			tasks.PUT("/:id", h.updateTask)
		}
		results := api.Group("/results")
		{
			results.POST("/", h.createResult)
			results.GET("/", h.getAllResults)
			results.GET("/task/:task_id", h.getResultsByTaskID)
			results.GET("/agent/:agent_id", h.getLatestResultsByAgent)
			results.DELETE("/:id", h.deleteResult)
		}
		metrics := api.Group("/metrics")
		{
			metrics.GET("/", h.getMetrics)
		}
	}

	return router
}
