package service

import (
	"delivery/internal/model"
	"delivery/internal/repository"
)

type DeliveryService struct {
	repo repository.Delivery
}

func NewDeliveryService(repo repository.Delivery) *DeliveryService {
	return &DeliveryService{repo: repo}
}

func (s *DeliveryService) Create(d model.Delivery) (int, error) {
	return s.repo.Create(d)
}

func (s *DeliveryService) GetById(id int) (model.Delivery, error) {
	return s.repo.GetById(id)
}

func (s *DeliveryService) GetAll(limit int, offset int) ([]model.Delivery, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *DeliveryService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *DeliveryService) Update(d model.Delivery) error {
	return s.repo.Update(d)
}
