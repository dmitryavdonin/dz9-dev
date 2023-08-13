package repository

import (
	"context"
	"net/http"
	"order/internal/model"
	"order/internal/repository/store/dto"
	"order/internal/service/adapters/store"
)

type StoreApi struct {
	baseUrl string
}

func NewStoreApi(baseUrl string) *StoreApi {
	return &StoreApi{baseUrl: baseUrl}
}

func (r *StoreApi) PlaceOrderInStore(ctx context.Context, storeOrder store.StoreOrderInfo) (statusResponse model.StatusResponse, err error) {
	request := &dto.PlaceOrderInStoreRequest{
		OrderId:  storeOrder.OrderId,
		BookId:   storeOrder.BookId,
		Quantity: storeOrder.Quantity,
	}

	response := &dto.PlaceOrderInStoreResponse{}

	_, err = sendRequest(r.baseUrl+"/order", http.MethodPost, "application/json", request, response)
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

func (r *StoreApi) CancelOrderInStore(ctx context.Context, orderId int, reason string) error {
	request := &dto.CancelOrderInStoreRequest{
		OrderId: orderId,
		Reason:  reason,
	}

	response := &dto.CancelOrderInStoreResponse{}

	var _, err = sendRequest(r.baseUrl+"/order/cancel", http.MethodPost, "application/json", request, response)
	return err
}
