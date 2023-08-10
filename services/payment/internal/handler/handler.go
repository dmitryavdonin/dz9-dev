package handler

import (
	"payment/internal/service"

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

	api := router.Group("/payment")
	{
		api.POST("/", h.createPayment)
		api.POST("/cancel/:id", h.cancelPayment)
		api.GET("/:id", h.getById)
		api.GET("/", h.getAll)
		api.DELETE("/:id", h.deletePayment)

	}

	return router
}
