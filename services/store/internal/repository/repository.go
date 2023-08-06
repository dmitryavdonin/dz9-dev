package repository

import (
	"store/internal/model"

	"gorm.io/gorm"
)

type StoreOrder interface {
	Create(input model.StoreOrder) (int, error)
	GetByOrderId(id int) (model.StoreOrder, error)
	DeleteByOrderId(id int) error
}

type Product interface {
	Create(order model.Product) (int, error)
	GetById(id int) (model.Product, error)
	Delete(id int) error
	Update(ind int, input model.Product) error
}

type Repository struct {
	StoreOrder
	Product
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		StoreOrder: NewStoreOrderPostgres(db),
		Product:    NewProductPostgres(db),
	}
}
