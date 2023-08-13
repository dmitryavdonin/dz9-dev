package repository

import (
	"context"
	"net/http"
	"order/internal/model"
	"order/internal/repository/payment/dto"
	"order/internal/service/adapters/payment"
)

type PaymentApi struct {
	baseUrl string
}

func NewPaymentApi(baseUrl string) *BookApi {
	return &BookApi{baseUrl: baseUrl}
}

func (r *BookApi) DoPayment(ctx context.Context, info payment.PaymentInfo) (statusResponse model.StatusResponse, err error) {
	request := &dto.DoPaymentRequest{
		OrderId: info.OrderId,
		UserId:  info.UserId,
		Money:   info.Money,
	}

	response := &dto.DoPaymentResponse{}

	_, err = sendRequest(r.baseUrl, http.MethodPost, "application/json", request, response)
	if err != nil {

		return model.StatusResponse{
			Status: "failed",
			Reason: err.Error(),
		}, err
	}

	return model.StatusResponse{
		Status: response.Status,
		Reason: response.Reason,
	}, nil
}
