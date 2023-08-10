package service

import (
	"context"
	"payment/internal/model"
	"payment/internal/repository"
)

type Payment interface {
	Create(pay model.Payment) (int, error)
	GetById(orderId int) (model.Payment, error)
	GetAll(limit int, offset int) ([]model.Payment, error)
	Delete(orderId int) error
	Update(pay model.Payment) error
}

type User interface {
	GetBalance(ctx context.Context, user_id int) (model.UserBalance, error)
	UpdateBalance(ctx context.Context, ub model.UserBalance) error
}

type Service struct {
	Payment
	User
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Payment: NewPaymentService(repos.Payment),
		User:    NewUserService(repos.User),
	}
}
