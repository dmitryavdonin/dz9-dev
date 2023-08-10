package service

import (
	"context"
	"order/internal/model"
	"order/internal/service/adapters/store"

	"github.com/sirupsen/logrus"
)

type SagaService struct {
	repo store.StoreApi
}

func NewSagaService(repo store.StoreApi) *SagaService {
	return &SagaService{repo: repo}
}

func (s *SagaService) CreateOrder(ctx context.Context, order model.Order) *model.StatusResponse {

	var response = model.StatusResponse{}

	var storeOrderInfo = store.StoreOrderInfo{
		OrderId:  order.ID,
		BookId:   order.BookId,
		Quantity: order.Quantity,
	}
	result, err := s.repo.CreateStoreOrder(ctx, storeOrderInfo)

	if err != nil {
		logrus.Errorf("Cannot create order with id = %d in store, error = %s", order.ID, err.Error())
		response.Status = "failed"
		response.Reason = err.Error()

	} else {

		response.Status = result.Status
		response.Reason = result.Reason
	}

	return &response
}

func (s *SagaService) CreateStoreOrder(ctx context.Context, storeOrderInfo store.StoreOrderInfo) (*model.StatusResponse, error) {
	return s.repo.CreateStoreOrder(ctx, storeOrderInfo)
}
