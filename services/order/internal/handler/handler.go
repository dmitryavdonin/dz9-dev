package handler

import (
	"order/internal/service"

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

	api := router.Group("/order")
	{
		api.POST("/", h.createOrder)
		api.GET("/:id", h.getOrderById)
		api.DELETE("/:id", h.deleteOrder)

	}

	return router
}
