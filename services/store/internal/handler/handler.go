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
		order_api.POST("/cancel/:id", h.cancelStoreOrder)
		order_api.GET("/", h.getAllStoreOrders)
		order_api.GET("/:id", h.getStoreOrderById)
		order_api.DELETE("/:id", h.deleteStoreOrder)
	}

	product_api := router.Group("/store/book")
	{
		product_api.POST("/", h.addBookToStore)
		product_api.GET("/", h.getAllStoreBooks)
		product_api.GET("/:id", h.getStoreBookById)
		product_api.DELETE("/:id", h.deleteStoreBook)
		product_api.PUT("/:id", h.updateStoreBook)
	}

	return router
}
