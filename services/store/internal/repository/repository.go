package repository

import (
	"store/internal/model"

	"gorm.io/gorm"
)

type StoreOrder interface {
	Create(input model.StoreOrder) (int, error)
	GetById(id int) (model.StoreOrder, error)
	GetAll(limit int, offset int) ([]model.StoreOrder, error)
	Delete(id int) error
	Update(ind int, input model.StoreOrder) error
	AlreadyExists(id int) bool
}

type StoreBook interface {
	Create(order model.StoreBook) (int, error)
	GetById(id int) (model.StoreBook, error)
	GetAll(limit int, offset int) ([]model.StoreBook, error)
	Delete(id int) error
	Update(ind int, input model.StoreBook) error
	AlreadyExists(id int) bool
}

type Repository struct {
	StoreOrder
	StoreBook
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		StoreOrder: NewStoreOrderPostgres(db),
		StoreBook:  NewStoreBookPostgres(db),
	}
}
