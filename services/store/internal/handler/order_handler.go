package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"store/internal/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createStoreOrder(c *gin.Context) {
	var input model.StoreOrder
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	// check if store order for this order_id already exists
	if h.service.StoreOrder.AlreadyExists(input.OrderId) {
		c.JSON(http.StatusConflict, StatusResponse{
			Status: "failed",
			Reason: fmt.Sprintf("Order with id = %d already exists", input.OrderId),
		})
		return
	}

	storeBook, err := h.service.StoreBook.GetById(input.BookId)
	if err != nil {
		// cannot get book in store info
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	oldInStock := storeBook.InStock

	now := time.Now()

	order := model.StoreOrder{
		OrderId:    input.OrderId,
		BookId:     input.BookId,
		Quantity:   input.Quantity,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	if storeBook.InStock >= order.Quantity {
		storeBook.InStock -= order.Quantity
		// update in_stock value
		err = h.service.StoreBook.Update(storeBook.BookId, storeBook)
		if err != nil {
			// return error with books
			c.JSON(http.StatusInternalServerError, StatusResponse{
				Status: "failed",
				Reason: err.Error(),
			})
			return
		}
		order.Status = "success"
		order.Reason = ""
	} else {
		order.Status = "failed"
		order.Reason = "Not enough amout"
	}

	// create order
	id, err := h.service.StoreOrder.Create(order)
	if err != nil {
		// restore old in_stock value
		storeBook.InStock = oldInStock

		h.service.StoreBook.Update(storeBook.BookId, storeBook)

		// return error with order
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	var reason string = order.Reason
	if reason == "" {
		reason = strconv.Itoa(id)
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: order.Status,
		Reason: reason,
	})
}

// Cancel order
func (h *Handler) cancelStoreOrder(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	// get order
	order, err := h.service.StoreOrder.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	// check order status
	// only success orders can be canceled
	if order.Status == "failed" {
		c.JSON(http.StatusOK, StatusResponse{
			Status: "failed",
			Reason: "Failed order cannot be canceled",
		})
		return
	}
	if order.Status == "canceled" {
		c.JSON(http.StatusOK, StatusResponse{
			Status: "success",
			Reason: "canceled",
		})
		return
	}

	// get store book
	book, err := h.service.StoreBook.GetById(order.BookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	// return the amount of books to the store
	book.InStock += order.Quantity

	err = h.service.StoreBook.Update(book.BookId, book)

	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	// update order status to canceled
	order.Status = "canceled"

	err = h.service.StoreOrder.Update(order.OrderId, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "success",
		Reason: "canceled",
	})
}

func (h *Handler) getStoreOrderById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	order, err := h.service.StoreOrder.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (h *Handler) deleteStoreOrder(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	err = h.service.StoreOrder.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "success",
		Reason: strconv.Itoa(id),
	})
}

func (h *Handler) getAllStoreOrders(c *gin.Context) {

	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var items []model.StoreOrder
	items, err := h.service.StoreOrder.GetAll(intLimit, offset)
	if err != nil {
		c.JSON(http.StatusBadGateway,
			StatusResponse{
				Status: "failed",
				Reason: err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(items), "data": items})
}
