package service

import (
	"order/internal/model"
	"order/internal/repository"
)

type Order interface {
	Create(order model.Order) (int, error)
	GetById(orderId int) (model.Order, error)
	Delete(orderId int) error
}

type Service struct {
	Order
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
	}
}
