package repository

import (
	"context"
	"order/internal/model"
	"order/internal/service/adapters/store"

	//"order/internal/service/adapters/repository"

	"gorm.io/gorm"
)

type Order interface {
	Create(order model.Order) (int, error)
	GetById(orderId int) (model.Order, error)
	Delete(orderId int) error
	Update(orderId int, input model.Order) error
}

type Store interface {
	CreateStoreOrder(ctx context.Context, storeOrder store.StoreOrderInfo) (statusResponse *model.StatusResponse, err error)
}

type Repository struct {
	Order
	Store
}

func NewRepository(db *gorm.DB, storeApiUri string) *Repository {
	return &Repository{
		Order: NewOrderPostgres(db),
		Store: NewStoreApi(storeApiUri),
	}
}
