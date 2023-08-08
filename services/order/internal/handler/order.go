package handler

import (
	"net/http"
	"strconv"
	"time"

	"order/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func (h *Handler) createOrderSaga(c *gin.Context, order_id int) {

	order, err := h.services.Order.GetById(order_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{Status: "failed", Reason: err.Error()})
		return
	}

	result := h.services.Saga.CreateOrder(c, order)

	order.Status = result.Status
	order.Reason = result.Reason

	err = h.services.Order.Update(order_id, order)
	if err != nil {
		logrus.Errorf("error occured on server shutting down: %s", err.Error())
	}

}

func (h *Handler) createOrder(c *gin.Context) {
	var input model.NewOrder
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{Status: "failed", Reason: err.Error()})
		return
	}

	var date, err = time.Parse("2006-01-02", input.DeliveryDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{Status: "failed", Reason: err.Error()})
		return
	}

	now := time.Now()

	order := model.Order{
		BookId:          input.BookId,
		Quantity:        input.Quantity,
		DeliveryAddress: input.DeliveryAddress,
		DeliveryDate:    date,
		Status:          "pending",
		CreatedAt:       now,
		ModifiedAt:      now,
	}

	id, err := h.services.Order.Create(order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{Status: "failed", Reason: err.Error()})
		return
	}

	c.JSON(http.StatusOK, StatusResponse{Status: "pending", Reason: strconv.Itoa(id)})

	go h.createOrderSaga(c, id)
}

func (h *Handler) getOrderById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order, err := h.services.Order.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, order)
}

func (h *Handler) deleteOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.Order.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
