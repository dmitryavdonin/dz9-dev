package service

import (
	"store/internal/model"
	"store/internal/repository"
)

type StoreOrder interface {
	Create(so model.StoreOrder) (int, error)
	GetByOrderId(id int) (model.StoreOrder, error)
	DeleteByOrderId(id int) error
}

type Product interface {
	Create(so model.Product) (int, error)
	GetById(id int) (model.Product, error)
	Delete(id int) error
	Update(id int, input model.Product) error
}

type Service struct {
	StoreOrder
	Product
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		StoreOrder: NewStoreOrderService(repos.StoreOrder),
		Product:    NewProductService(repos.Product),
	}
}
