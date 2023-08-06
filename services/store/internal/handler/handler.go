package handler

import (
	"store/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{service: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	order_api := router.Group("/store/order")
	{
		order_api.POST("/", h.createStoreOrder)
		order_api.GET("/:id", h.getStoreOrderByOrderId)
		order_api.DELETE("/:id", h.deleteStoreOrderByOrderId)
	}

	product_api := router.Group("/store/product")
	{
		product_api.POST("/", h.createProduct)
		product_api.GET("/:id", h.getProductById)
		product_api.DELETE("/:id", h.deleteProduct)
		product_api.PUT("/:id", h.updateProduct)
	}

	return router
}
