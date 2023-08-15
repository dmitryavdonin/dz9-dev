package repository

import (
	"context"
	"net/http"
	"order/internal/model"
	"order/internal/repository/payment/dto"
	"order/internal/service/adapters/payment"

	"github.com/sirupsen/logrus"
)

type PaymentApi struct {
	baseUrl string
}

func NewPaymentApi(baseUrl string) *PaymentApi {
	return &PaymentApi{baseUrl: baseUrl}
}

// do payment
func (r *PaymentApi) DoPayment(ctx context.Context, info payment.PaymentInfo) (statusResponse model.StatusResponse, err error) {

	logrus.Printf("DoPaymnet: BEGIN Try to create payment for order_d = %d, user_id = %d, money = %d",
		info.OrderId, info.UserId, info.Money)

	request := &dto.DoPaymentRequest{
		OrderId: info.OrderId,
		UserId:  info.UserId,
		Money:   info.Money,
	}

	response := &dto.DoPaymentResponse{}

	_, err = sendRequest(r.baseUrl, http.MethodPost, "application/json", request, response)
	if err != nil {
		logrus.Errorf("DoPaymnet: Cannot create payment for order_d = %d, user_id = %d, money = %d, error = %s",
			info.OrderId, info.UserId, info.Money, err.Error())
		return model.StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		}, err
	}

	logrus.Printf("DoPaymnet: END Payment created for order_d = %d, user_id = %d, money = %d, status = %s, reason = %",
		info.OrderId, info.UserId, info.Money, response.Status, response.Reason)

	return model.StatusResponse{
		Status: response.Status,
		Reason: response.Reason,
	}, nil
}

// cancel payment
func (r *PaymentApi) CancelPayment(ctx context.Context, orderId int, reason string) error {

	logrus.Printf("CancelPayment(): BEGIN Try to cancel payment for order_d = %d, reason = %s",
		orderId, reason)

	request := &dto.CancelPaymentRequest{
		OrderId: orderId,
		Reason:  reason,
	}

	response := &dto.CancelPaymentResponse{}

	var _, err = sendRequest(r.baseUrl+"/cancel", http.MethodPost, "application/json", request, response)
	if err != nil {
		logrus.Errorf("CancelPayment(): Cannot cancel payment for order_d = %d, error = %s",
			orderId, err.Error())
	} else {
		logrus.Printf("CancelPayment(): END payment canceled for order_d = %d, reason = %s",
			orderId, reason)
	}
	return err
}
