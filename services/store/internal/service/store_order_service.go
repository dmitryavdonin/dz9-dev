package service

import (
	"store/internal/model"
	"store/internal/repository"
)

type StoreOrderService struct {
	repo repository.StoreOrder
}

func NewStoreOrderService(repo repository.StoreOrder) *StoreOrderService {
	return &StoreOrderService{repo: repo}
}

func (s *StoreOrderService) Create(order model.StoreOrder) (int, error) {
	return s.repo.Create(order)
}

func (s *StoreOrderService) GetByOrderId(order_id int) (model.StoreOrder, error) {
	return s.repo.GetByOrderId(order_id)
}

func (s *StoreOrderService) DeleteByOrderId(order_id int) error {
	return s.repo.DeleteByOrderId(order_id)
}
