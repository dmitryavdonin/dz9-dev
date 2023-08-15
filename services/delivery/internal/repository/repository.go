package repository

import (
	"delivery/internal/model"

	"gorm.io/gorm"
)

type Delivery interface {
	Create(order model.Delivery) (int, error)
	GetById(orderId int) (model.Delivery, error)
	GetAll(limit int, offset int) ([]model.Delivery, error)
	Delete(orderId int) error
	Update(input model.Delivery) error
}

type Repository struct {
	Delivery
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Delivery: NewDeliveryPostgres(db),
	}
}
