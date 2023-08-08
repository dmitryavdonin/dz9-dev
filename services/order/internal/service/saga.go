package service

import (
	"context"
	"order/internal/model"
	"order/internal/service/adapters/store"
)

type SagaService struct {
	repo store.StoreApi
}

func NewSagaService(repo store.StoreApi) *SagaService {
	return &SagaService{repo: repo}
}

func (s *SagaService) CreateOrder(ctx context.Context, order model.Order) *model.StatusResponse {

	var storeOrderInfo = store.StoreOrderInfo{
		OrderId:  order.ID,
		BookId:   order.BookId,
		Quantity: order.Quantity,
	}
	result, err := s.repo.CreateStoreOrder(ctx, storeOrderInfo)

	if err != nil {

	}

	if result.Status == "failed" {

	}

	return &model.StatusResponse{
		Status: result.Status,
		Reason: result.Reason,
	}
}

func (s *SagaService) CreateStoreOrder(ctx context.Context, storeOrderInfo store.StoreOrderInfo) (*model.StatusResponse, error) {
	return s.repo.CreateStoreOrder(ctx, storeOrderInfo)
}
