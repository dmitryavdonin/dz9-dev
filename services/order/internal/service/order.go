package service

import (
	"order/internal/model"
	"order/internal/repository"
)

type OrderService struct {
	repo repository.Order
}

func NewOrderService(repo repository.Order) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) Create(order model.Order) (int, error) {
	return s.repo.Create(order)
}

func (s *OrderService) GetById(id int) (model.Order, error) {
	return s.repo.GetById(id)
}

func (s *OrderService) Delete(id int) error {
	return s.repo.Delete(id)
}