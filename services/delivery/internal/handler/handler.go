package handler

import (
	"delivery/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/delivery")
	{
		api.POST("/", h.createDelivery)
		api.POST("/cancel/:id", h.cancelDelivery)
		api.GET("/:id", h.getById)
		api.GET("/", h.getAll)
		api.DELETE("/:id", h.deleteDelivery)
	}

	return router
}
