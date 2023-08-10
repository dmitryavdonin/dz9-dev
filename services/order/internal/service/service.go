package service

import (
	"context"
	"order/internal/model"
	"order/internal/repository"
)

type Order interface {
	Create(order model.Order) (int, error)
	GetById(orderId int) (model.Order, error)
	GetAll(limit int, offset int) ([]model.Order, error)
	Delete(orderId int) error
	Update(id int, order model.Order) error
}

type Saga interface {
	CreateOrder(ctx context.Context, order model.Order) *model.StatusResponse
}

type Service struct {
	Order
	Saga
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Order: NewOrderService(repos.Order),
		Saga:  NewSagaService(repos.Store),
	}
}
