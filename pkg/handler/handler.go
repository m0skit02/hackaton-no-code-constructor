package handler

import (
	"hackaton-no-code-constructor/pkg/service"

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
		blockTypes := api.Group("/blockTypes")
		{
			blockTypes.POST("/", h.createBlockType)      // Создать блок
			blockTypes.GET("/", h.getListBlockTypes)     //Получить список блоков
			blockTypes.GET("/:id", h.getByIDBlockType)   //Получить блок по ID
			blockTypes.PUT("/:id", h.updateBlockType)    //Обновить блок по ID
			blockTypes.DELETE("/:id", h.deleteBlockType) //Удалить блок по ID
		}
	}
	return router
}
