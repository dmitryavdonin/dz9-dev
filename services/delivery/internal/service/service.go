package service

import (
	"delivery/internal/model"
	"delivery/internal/repository"
)

type Delivery interface {
	Create(d model.Delivery) (int, error)
	GetById(orderId int) (model.Delivery, error)
	GetAll(limit int, offset int) ([]model.Delivery, error)
	Delete(orderId int) error
	Update(d model.Delivery) error
}

type Service struct {
	Delivery
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Delivery: NewDeliveryService(repos.Delivery),
	}
}
