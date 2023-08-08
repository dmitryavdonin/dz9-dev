package service

import (
	"store/internal/model"
	"store/internal/repository"
)

type StoreOrder interface {
	Create(so model.StoreOrder) (int, error)
	GetById(order_id int) (model.StoreOrder, error)
	GetAll(limit int, offset int) ([]model.StoreOrder, error)
	Delete(order_id int) error
	Update(order_id int, input model.StoreOrder) error
	AlreadyExists(order_id int) bool
}

type StoreBook interface {
	Create(so model.StoreBook) (int, error)
	GetById(id int) (model.StoreBook, error)
	GetAll(limit int, offset int) ([]model.StoreBook, error)
	Delete(id int) error
	Update(id int, input model.StoreBook) error
	AlreadyExists(id int) bool
}

type Service struct {
	StoreOrder
	StoreBook
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		StoreOrder: NewStoreOrderService(repos.StoreOrder),
		StoreBook:  NewStoreBookService(repos.StoreBook),
	}
}
