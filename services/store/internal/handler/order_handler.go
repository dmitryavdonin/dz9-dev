package handler

import (
	"net/http"
	"strconv"
	"time"

	"store/internal/model"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// create a new order in the store
func (h *Handler) createStoreOrder(c *gin.Context) {
	logrus.Printf("createStoreOrder(): BEGIN")
	var input model.StoreOrder
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("createStoreOrder(): Cannot parse input, error = %s", err.Error())
		return
	}

	// at first try to find the required book in the store
	logrus.Printf("createStoreOrder(): Try to get book in store by book_id = %d", input.BookId)
	book, err := h.service.StoreBook.GetById(input.BookId)
	if err != nil {
		// cannot get book in store
		c.JSON(http.StatusNotFound, StatusResponse{
			Status: "failed",
			Reason: "The required book not found",
		})
		logrus.Errorf("createStoreOrder(): Cannot get book_id = %d, error = %s", input.BookId, err.Error())
		return
	}

	// check if an order with such ID already exist in the store
	logrus.Printf("createStoreOrder(): Check if order_id = %d already exist", input.OrderId)
	existentOrder, err := h.service.StoreOrder.GetById(input.OrderId)
	if err == nil {
		logrus.Printf("createStoreOrder(): order_id = %d already exist, so try to update the order in store", input.OrderId)
		h.updateExistentStoreOrder(c, input, existentOrder, book)
	} else {
		logrus.Printf("createStoreOrder(): order_id = %d not already exist, so try to create a new the order in store", input.OrderId)
		h.createNewStoreOrder(c, input, book)
	}
}

// update book in store if the amout of books is enough for the order
func (h *Handler) updateBookInStockIfEnough(quantity int, book model.StoreBook) (bool, error) {
	// check if we have enough books in the store
	logrus.Printf("updateBookInStockIfEnough(): Check if we have enough books in the store: in_stock = %d, quantity = %d", book.InStock, quantity)
	if book.InStock >= quantity {
		book.InStock -= quantity
		// update in_stock value
		logrus.Printf("updateBookInStockIfEnough(): Try to update in_stock = %d, quantity = %d", book.InStock, quantity)
		err := h.service.StoreBook.Update(book.BookId, book)
		if err != nil {
			logrus.Errorf("updateBookInStockIfEnough(): Cannot update in_stock value, error= %s", err.Error())
			return false, err
		}
		return true, nil
	}
	logrus.Print("updateBookInStockIfEnough(): Not enough books in the store")
	return false, nil
}

// create a new store order
func (h *Handler) createNewStoreOrder(c *gin.Context,
	input model.StoreOrder, book model.StoreBook) {

	logrus.Print("createNewStoreOrder(): BEGIN")

	now := time.Now()

	newOrder := model.StoreOrder{
		OrderId:    input.OrderId,
		BookId:     input.BookId,
		Quantity:   input.Quantity,
		CreatedAt:  now,
		ModifiedAt: now,
	}

	// just remeber the old vaule to restore it if any issue happens
	oldInStock := book.InStock

	logrus.Printf("createNewStoreOrder(): Try to update book in stock value if enough oldInStock = %d, quantity = %d", oldInStock, newOrder.Quantity)
	result, err := h.updateBookInStockIfEnough(newOrder.Quantity, book)
	if err != nil {
		// return error with books
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})

		logrus.Errorf("createNewStoreOrder(): Cannot update in_stock value, error= %s", err.Error())
	}

	if result {
		newOrder.Status = "success"
		newOrder.Reason = ""
	} else {
		newOrder.Status = "failed"
		newOrder.Reason = "Not enough books in the store"
	}

	// create order
	logrus.Printf("createNewStoreOrder(): Try to create new order book_id = %d, quantity = %d, status = %s, reason = %s",
		newOrder.BookId, newOrder.Quantity, newOrder.Status, newOrder.Reason)

	id, err := h.service.StoreOrder.Create(newOrder)
	if err != nil {
		logrus.Errorf("createNewStoreOrder(): Cannot create a new order book_id = %d, quantity = %d, error= %s",
			newOrder.BookId, newOrder.Quantity, err.Error())

		// restore old in_stock value
		logrus.Printf("createNewStoreOrder(): Try restore the books amount int the srore oldInStock = %d, book_id = %d, quantity = %d",
			oldInStock, newOrder.BookId, newOrder.Quantity)

		book.InStock = oldInStock

		if err1 := h.service.StoreBook.Update(book.BookId, book); err1 != nil {
			logrus.Errorf("createNewStoreOrder(): Cannot restore the oldInStock = %d, error= %s",
				oldInStock, err1.Error())
		}

		// return error with order
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		return
	}

	newOrder.ID = id

	c.JSON(http.StatusOK, newOrder)

	logrus.Printf("createNewStoreOrder(): END order created with order_id = %d, status = %s, reason = %s",
		newOrder.ID, newOrder.Status, newOrder.Reason)
}

// update existent store order
func (h *Handler) updateExistentStoreOrder(c *gin.Context,
	input model.StoreOrder, existentOrder *model.StoreOrder, book model.StoreBook) {
	// the order already exists and already successfuy processed
	// and nothing changed, then just do nothing and return existant order
	logrus.Printf("updateExistentStoreOrder(): BEGIN order_id = %d, existent order status = %s",
		input.OrderId, existentOrder.Status)

	if existentOrder.Status == "success" {
		if existentOrder.BookId == input.BookId &&
			existentOrder.Quantity == input.Quantity {
			c.JSON(http.StatusOK, existentOrder)
			logrus.Printf("updateExistentStoreOrder(): No need to update because existen order_id = %d already in success status",
				input.OrderId)
			return
		} else {
			// that's not possible to change any parameters of the order which was successfuly completed
			c.JSON(http.StatusBadRequest, StatusResponse{
				Status: "failed",
				Reason: "Cannot change parameters of the completed order",
			})
			logrus.Printf("updateExistentStoreOrder(): Cannot change parameters of the completed order_id = %d",
				input.OrderId)
			return
		}
	} else {
		// so, the order exists but processing was not successful (status is failed or canceled)
		// we just try to process such and order once again
		// just remeber the old vaule to restore it if any issue happens
		oldInStock := book.InStock

		logrus.Printf("updateExistentStoreOrder(): Try to update book in stock if enough, order_id = %d, quantity = %d, book_id = %d",
			existentOrder.OrderId, existentOrder.Quantity, book.BookId)

		bookStoreUpdated, err := h.updateBookInStockIfEnough(existentOrder.Quantity, book)
		if err != nil {
			// return error with books
			c.JSON(http.StatusInternalServerError, StatusResponse{
				Status: "failed",
				Reason: err.Error(),
			})

			logrus.Errorf("updateExistentStoreOrder(): Cannot update book in stock if enough, order_id = %d, error = %s",
				existentOrder.OrderId, err.Error())
		}

		if bookStoreUpdated {
			existentOrder.Status = "success"
			existentOrder.Reason = ""
		} else {
			existentOrder.Status = "failed"
			existentOrder.Reason = "Not enough books in the store"
		}

		existentOrder.ModifiedAt = time.Now()

		// update existent order
		logrus.Printf("updateExistentStoreOrder(): Try to update existent order_id = %d, quantity = %d, book_id = %d",
			existentOrder.OrderId, existentOrder.Quantity, book.BookId)

		err = h.service.StoreOrder.Update(input.OrderId, existentOrder)
		if err != nil {
			// restore old in_stock value
			book.InStock = oldInStock
			h.service.StoreBook.Update(book.BookId, book)
			// return error with order
			c.JSON(http.StatusInternalServerError, StatusResponse{
				Status: "failed",
				Reason: err.Error(),
			})

			logrus.Errorf("updateExistentStoreOrder(): Cannot update existent order_id = %d, error = %s",
				existentOrder.OrderId, err.Error())

			return
		}
		c.JSON(http.StatusOK, existentOrder)

		// update existent order
		logrus.Printf("updateExistentStoreOrder(): END order_id = %d, quantity = %d, book_id = %d",
			existentOrder.OrderId, existentOrder.Quantity, book.BookId)
	}
}

// Cancel order
func (h *Handler) cancelStoreOrder(c *gin.Context) {

	logrus.Print("cancelStoreOrder(): BEGIN")

	var input model.CancelOrder
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelStoreOrder(): Cannot parse input, error = %s", err.Error())
		return
	}

	// get order
	logrus.Printf("cancelStoreOrder(): Try to get store order by order_id = %d", input.OrderId)
	order, err := h.service.StoreOrder.GetById(input.OrderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})

		logrus.Errorf("cancelStoreOrder(): Cannot get store order by order_id = %d, error = %s", input.OrderId, err.Error())

		return
	}

	// check order status
	// only success orders can be canceled
	if order.Status == "failed" {
		c.JSON(http.StatusOK, StatusResponse{
			Status: "failed",
			Reason: "Failed order cannot be canceled",
		})

		logrus.Printf("cancelStoreOrder(): Cannot cancel the order order_id = %d because store order status is failed, only success store orders can be canceled", input.OrderId)

		return
	}
	if order.Status == "canceled" {
		c.JSON(http.StatusOK, StatusResponse{
			Status: "success",
			Reason: "canceled",
		})
		logrus.Printf("cancelStoreOrder(): Store order_id = %d is already canceled", input.OrderId)
		return
	}

	// get store book
	logrus.Printf("cancelStoreOrder(): Try to get the book with book_id = %d, order_id = %d", order.BookId, input.OrderId)
	book, err := h.service.StoreBook.GetById(order.BookId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelStoreOrder(): Cannot get book with book_id = %d, order_id = %d, error = %s", order.BookId, input.OrderId, err.Error())
		return
	}

	// return the amount of books to the store
	logrus.Printf("cancelStoreOrder(): Try to update the book with book_id = %d, order_id = %d, old in_stock = %d, new in_stock = %d", order.BookId, input.OrderId, book.InStock, order.Quantity)
	book.InStock += order.Quantity

	if err = h.service.StoreBook.Update(book.BookId, book); err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelStoreOrder(): Cannot update the book with book_id = %d, order_id = %d, error = %s", order.BookId, input.OrderId, err.Error())
		return
	}

	// update order status to canceled
	order.Status = "canceled"
	order.Reason = input.Reason
	order.ModifiedAt = time.Now()

	logrus.Printf("cancelStoreOrder(): Try to update the store order order_id = %d with cancel status and reason = %s ", input.OrderId, order.Reason)
	err = h.service.StoreOrder.Update(order.OrderId, order)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelStoreOrder(): Cannot update the store order order_id = %d, error = %s", input.OrderId, err.Error())
		return
	}

	c.JSON(http.StatusOK, StatusResponse{
		Status: "success",
		Reason: "canceled",
	})

	logrus.Printf("cancelStoreOrder(): END order_id = %d, status = success, reason = canceled", input.OrderId)
}

// get store order by order_id
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
