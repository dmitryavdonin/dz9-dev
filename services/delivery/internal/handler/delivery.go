package handler

import (
	"delivery/internal/model"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// create payment
func (h *Handler) createDelivery(c *gin.Context) {
	logrus.Info("CreateDelivery(): BEGIN")
	var input model.NewDelivery
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("CreateDelivery(): Cannot parse input, error = %s", err.Error())
		return
	}

	var date, err = time.Parse("2006-01-02", input.DeliveryDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("CreateDelivery(): Cannot parse delivery date, error = %s", err.Error())
		return
	}

	now := time.Now()

	d := model.Delivery{
		UserId:          input.UserId,
		OrderId:         input.OrderId,
		DeliveryAddress: input.DeliveryAddress,
		DeliveryDate:    date,
		Status:          "pending",
		CreatedAt:       now,
		ModifiedAt:      now,
	}

	id, err := h.services.Delivery.Create(d)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("CreateDelivery(): Cannot create delivery, error = %s", err.Error())
		return
	}

	d.ID = id

	logrus.Printf("CreateDelivery(): Delivery for order_id = %d created with status pending", d.OrderId)

	// check delivery address
	wrongAddress := "Some strange address"

	if d.DeliveryAddress == wrongAddress {
		d.Status = "failed"
		d.Reason = "Address is out of delivery area"
		d.ModifiedAt = time.Now()
		h.services.Delivery.Update(d)
		logrus.Printf("CreateDelivery(): Cannot do the delivery, status = %s, reason = %s, order_id = %d", d.Status, d.Reason, d.OrderId)
		c.JSON(http.StatusOK, d)
		return
	}

	if d.DeliveryDate.Before(time.Now()) {
		d.Status = "failed"
		d.Reason = "No couriers available for this date"
		d.ModifiedAt = time.Now()
		h.services.Delivery.Update(d)
		logrus.Printf("CreateDelivery(): Cannot do the delivery, status = %s, reason = %s, order_id = %d", d.Status, d.Reason, d.OrderId)
		c.JSON(http.StatusOK, d)
		return
	}

	d.Status = "success"
	d.Reason = ""
	d.ModifiedAt = time.Now()
	c.JSON(http.StatusOK, d)

	if err := h.services.Delivery.Update(d); err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("CreateDelivery(): Cannot create delivery, error = %s", err.Error())
		return
	}

	logrus.Printf("CreateDelivery(): END Delivery created successfuly for order_id = %d", d.OrderId)
}

// get delivery by order_id
func (h *Handler) getById(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	d, err := h.services.Delivery.GetById(id)
	if err != nil {
		if err.Error() == "record not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		return
	}

	c.JSON(http.StatusOK, d)
}

// get all records
func (h *Handler) getAll(c *gin.Context) {

	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var items []model.Delivery
	items, err := h.services.Delivery.GetAll(intLimit, offset)
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
func (h *Handler) deleteDelivery(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.services.Delivery.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "success",
	})
}

// Cancel delivery
func (h *Handler) cancelDelivery(c *gin.Context) {
	logrus.Printf("cancelDelivery(): BEGIN")

	var input model.CancelDelivery
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, StatusResponse{Status: "failed", Reason: err.Error()})
		logrus.Errorf("cancelDelivery(): Cannot parse input, error = %s", err.Error())
		return
	}

	// get delivery
	d, err := h.services.Delivery.GetById(input.OrderId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelDelivery(): Cannot get delivery for order_id = %d, error = %s", input.OrderId, err.Error())
		return
	}

	logrus.Printf("cancelDelivery(): delivery status = %s for order_id = %d", d.Status, input.OrderId)

	// check delivery status
	// only success delivery can be canceled
	if d.Status == "failed" {
		c.JSON(http.StatusOK, StatusResponse{
			Status: "failed",
			Reason: "Failed delivery cannot be canceled",
		})
		logrus.Errorf("cancelDelivery(): Failed delivery cannot be canceled, order_id = %d", input.OrderId)
		return
	}
	if d.Status == "canceled" {
		c.JSON(http.StatusOK, d)
		logrus.Printf("cancelDelivery(): Delivery already canceled, order_id = %d", input.OrderId)
		return
	}

	// update delivery status to canceled
	d.Status = "canceled"
	d.Reason = input.Reason
	d.ModifiedAt = time.Now()

	logrus.Printf("cancelDelivery(): Try to update delivery status to canceled, reason = %s, order_id = %d", input.Reason, input.OrderId)
	if err := h.services.Delivery.Update(d); err != nil {
		c.JSON(http.StatusInternalServerError, StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		})
		logrus.Errorf("cancelDelivery(): Cannot update delivery status with canceled, order_id = %d, error = %s", input.OrderId, err.Error())
		return
	}

	c.JSON(http.StatusOK, d)

	logrus.Printf("cancelDelivery(): END, order_id = %d", input.OrderId)
}
