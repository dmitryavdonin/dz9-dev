package repository

import (
	"context"
	"net/http"
	"order/internal/model"
	"order/internal/repository/delivery/dto"
	"order/internal/service/adapters/delivery"

	"github.com/sirupsen/logrus"
)

type DeliveryApi struct {
	baseUrl string
}

func NewDeliveryApi(baseUrl string) *DeliveryApi {
	return &DeliveryApi{baseUrl: baseUrl}
}

// do delivery
func (r *DeliveryApi) DoDelivery(ctx context.Context, info delivery.DeliveryInfo) (statusResponse model.StatusResponse, err error) {

	logrus.Printf("DoDelivery(): BEGIN Try to create delivery for order_d = %d, user_id = %d, delivery_address = %s, dleivery_date = %s",
		info.OrderId, info.UserId, info.DeliveryAddress, info.DeliveryDate)

	request := &dto.DoDeliveryRequest{
		OrderId:         info.OrderId,
		UserId:          info.UserId,
		DeliveryAddress: info.DeliveryAddress,
		DeliveryDate:    info.DeliveryDate.Format("2006-01-02"),
	}

	response := &dto.DoDeliveryResponse{}

	_, err = sendRequest(r.baseUrl, http.MethodPost, "application/json", request, response)
	if err != nil {
		logrus.Errorf("DoDelivery(): Cannot create delivery for order_d = %d, user_id = %d, error = %s",
			info.OrderId, info.UserId, err.Error())
		return model.StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		}, err
	}

	logrus.Printf("DoDelivery(): END Delivery created for order_d = %d, user_id = %d, status = %s, reason = %s",
		info.OrderId, info.UserId, response.Status, response.Reason)

	return model.StatusResponse{
		Status: response.Status,
		Reason: response.Reason,
	}, nil
}

// cancel delivery
func (r *DeliveryApi) CancelDelivery(ctx context.Context, orderId int, reason string) error {

	logrus.Printf("CancelDelivery(): BEGIN Try to cancel delivery for order_d = %d, reason = %s",
		orderId, reason)

	request := &dto.CancelDeliveryRequest{
		OrderId: orderId,
		Reason:  reason,
	}

	response := &dto.CancelDeliveryResponse{}

	var _, err = sendRequest(r.baseUrl+"/cancel", http.MethodPost, "application/json", request, response)
	if err != nil {
		logrus.Errorf("DoPaymnet: Cannot cancel delivery for order_d = %d, error = %s",
			orderId, err.Error())
	} else {
		logrus.Printf("DoPaymnet: END delivery canceled for order_d = %d, reason = %s",
			orderId, reason)
	}
	return err
}
