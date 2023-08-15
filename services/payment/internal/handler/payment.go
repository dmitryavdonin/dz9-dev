package handler

import (
	"net/http"
	"payment/internal/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// create payment
func (h *Handler) createPayment(c *gin.Context) {
	logrus.Info("CreatePayment(): BEGIN")
	var input model.NewPayment
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("CreatePayment(): Cannot parse input, error = %s", err.Error())
		return
	}

	now := time.Now()

	pay := model.Payment{
		UserId:     input.UserId,
		OrderId:    input.OrderId,
		Money:      input.Money,
		Status:     "pending",
		CreatedAt:  now,
		ModifiedAt: now,
	}

	id, err := h.services.Payment.Create(pay)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("CreatePayment(): Cannot create payment, error = %s", err.Error())
		return
	}

	pay.ID = id

	logrus.Printf("CreatePayment(): Payment for order_id = %d created with status pending", pay.OrderId)

	//get user ub
	logrus.Printf("CreatePayment(): Try to get user balance for user_id = %d", pay.UserId)
	balance, err := h.services.User.GetBalance(c, pay.UserId)
	if err != nil {
		pay.Status = "failed"
		pay.Reason = err.Error()
		pay.ModifiedAt = time.Now()
		h.services.Payment.Update(pay)
		c.JSON(http.StatusInternalServerError, pay)
		logrus.Errorf("CreatePayment(): Cannot get user balance, error = %s", err.Error())
		return
	}

	pay.ModifiedAt = time.Now()

	if balance >= pay.Money {
		balance -= pay.Money

		logrus.Printf("CreatePayment(): Try to update user balance = %d for user_id = %d", balance, pay.UserId)
		if err := h.services.User.UpdateBalance(c, pay.UserId, balance); err != nil {
			pay.Status = "failed"
			pay.Reason = err.Error()
			pay.ModifiedAt = time.Now()
			h.services.Payment.Update(pay)

			c.JSON(http.StatusInternalServerError, pay)

			logrus.Errorf("CreatePayment(): Cannot update user balance, error = %s", err.Error())

			return
		}
	} else {
		pay.Status = "failed"
		pay.Reason = "Not enough balance"
		pay.ModifiedAt = time.Now()

		if err = h.services.Payment.Update(pay); err != nil {
			pay.Status = "failed"
			pay.Reason = err.Error()
			c.JSON(http.StatusInternalServerError, pay)
			logrus.Errorf("CreatePayment(): Cannot update payment for order_id = %d, error = %s", pay.OrderId, err.Error())
			return
		}

		c.JSON(http.StatusOK, pay)

		logrus.Printf("CreatePayment(): Payment failed. Not enough balance for order_id = %d", pay.OrderId)

		return
	}
	pay.Status = "success"
	pay.Reason = ""
	pay.ModifiedAt = time.Now()
	c.JSON(http.StatusOK, pay)

	logrus.Printf("CreatePayment(): Try to update payment with status = %s, reason = %s, order_id = %d", pay.Status, pay.Reason, pay.OrderId)
	if err = h.services.Payment.Update(pay); err != nil {
		c.JSON(http.StatusInternalServerError, pay)
		logrus.Errorf("CreatePayment(): Cannot update payment for order_id = %d, error = %s", pay.OrderId, err.Error())
		return
	}

	logrus.Printf("CreatePayment(): Payment created successfuly for order_id = %d", pay.OrderId)
}

// get payment by order_id
func (h *Handler) getById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pay, err := h.services.Payment.GetById(id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, pay)
}

// get all records
func (h *Handler) getAll(c *gin.Context) {

	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var items []model.Payment
	items, err := h.services.Payment.GetAll(intLimit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError,
			StatusResponse{
				Status: "failed",
				Reason: err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(items), "data": items})
}

// delete payment
func (h *Handler) deletePayment(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.Payment.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "success",
	})
}

// cancel payment
// Cancel order
func (h *Handler) cancelPayment(c *gin.Context) {
	logrus.Printf("cancelPayment(): BEGIN")

	var input model.CancelPayment
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("cancelPayment(): Cannot parse input, error = %s", err.Error())
		return
	}

	// get payment
	pay, err := h.services.Payment.GetById(input.OrderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelPayment(): Cannot get payment for order_id = %d, error = %s", input.OrderId, err.Error())
		return
	}

	logrus.Printf("cancelPayment(): payment status = %s for order_id = %d", pay.Status, input.OrderId)

	// check pay status
	// only success payment can be canceled
	if pay.Status == "failed" {
		c.JSON(http.StatusOK, StatusResponse{
			Status: "failed",
			Reason: "Failed payment cannot be canceled",
		})
		logrus.Errorf("cancelPayment(): Failed payment cannot be canceled, order_id = %d", input.OrderId)
		return
	}
	if pay.Status == "canceled" {
		c.JSON(http.StatusOK, pay)
		logrus.Printf("cancelPayment(): Payment already canceled, order_id = %d", input.OrderId)
		return
	}

	// get user balance
	logrus.Printf("cancelPayment(): Try to get user balance, user_id = %d, order_id = %d", pay.UserId, input.OrderId)
	balance, err := h.services.User.GetBalance(c, pay.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelPayment(): Cannot get user balance user_id = %d, order_id = %d, error = %s", pay.UserId, input.OrderId, err.Error())
		return
	}

	// return the amount of money back to the user's balance
	balance += pay.Money
	logrus.Printf("cancelPayment(): Try to return the money = %d tot the user balance = %d, user_id = %d, order_id = %d", pay.Money, balance, pay.UserId, input.OrderId)
	if err := h.services.User.UpdateBalance(c, pay.UserId, balance); err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelPayment(): Cannot return the money to the user balance, user_id = %d, order_id = %d, error = %s", pay.UserId, input.OrderId, err.Error())
		return
	}

	// update order status to canceled
	pay.Status = "canceled"
	pay.Reason = input.Reason
	pay.ModifiedAt = time.Now()

	logrus.Printf("cancelPayment(): Try to update payment status with canceled, reason = %s, order_id = %d", input.Reason, input.OrderId)
	if err := h.services.Payment.Update(pay); err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelPayment(): Cannot update payment status with canceled, order_id = %d, error = %s", input.OrderId, err.Error())
		return
	}

	c.JSON(http.StatusOK, pay)

	logrus.Printf("cancelPayment(): END, order_id = %d", input.OrderId)
}
