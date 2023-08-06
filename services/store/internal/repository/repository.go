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

type StoreProduct interface {
	Create(order model.StoreProduct) (int, error)
	GetById(id int) (model.StoreProduct, error)
	GetByProductId(id int) (model.StoreProduct, error)
	Delete(id int) error
	Update(ind int, input model.StoreProduct) error
}

type Repository struct {
	StoreOrder
	StoreProduct
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		StoreOrder:   NewStoreOrderPostgres(db),
		StoreProduct: NewProductPostgres(db),
	}
}
