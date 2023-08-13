package repository

import (
	"context"
	"payment/internal/model"

	"gorm.io/gorm"
)

type Payment interface {
	Create(order model.Payment) (int, error)
	GetById(orderId int) (model.Payment, error)
	GetAll(limit int, offset int) ([]model.Payment, error)
	Delete(orderId int) error
	Update(input model.Payment) error
}

type User interface {
	GetBalance(ctx context.Context, user_id int) (int, error)
	UpdateBalance(ctx context.Context, user_id int, balance int) error
}

type Repository struct {
	Payment
	User
}

func NewRepository(db *gorm.DB, userApiUri string) *Repository {
	return &Repository{
		Payment: NewPaymentPostgres(db),
		User:    NewUserApi(userApiUri),
	}
}
