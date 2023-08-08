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

func (s *StoreOrderService) GetById(order_id int) (model.StoreOrder, error) {
	return s.repo.GetById(order_id)
}

func (s *StoreOrderService) GetAll(limit int, offset int) ([]model.StoreOrder, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *StoreOrderService) Delete(order_id int) error {
	return s.repo.Delete(order_id)
}

func (s *StoreOrderService) Update(order_id int, input model.StoreOrder) error {
	return s.repo.Update(order_id, input)
}

func (s *StoreOrderService) AlreadyExists(order_id int) bool {
	return s.repo.AlreadyExists(order_id)
}
