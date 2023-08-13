package repository

import (
	"context"
	"order/internal/model"
	"order/internal/service/adapters/payment"
	"order/internal/service/adapters/store"

	//"order/internal/service/adapters/repository"

	"gorm.io/gorm"
)

type Order interface {
	Create(order model.Order) (int, error)
	GetById(orderId int) (model.Order, error)
	GetAll(limit int, offset int) ([]model.Order, error)
	Delete(orderId int) error
	Update(orderId int, input model.Order) error
}

type Store interface {
	PlaceOrderInStore(ctx context.Context, storeOrder store.StoreOrderInfo) (statusResponse model.StatusResponse, err error)
	CancelOrderInStore(ctx context.Context, orderId int, reason string) error
}

type Payment interface {
	DoPayment(ctx context.Context, info payment.PaymentInfo) (statusResponse model.StatusResponse, err error)
}

type Book interface {
	GetBookPrice(ctx context.Context, bookId int) (bookPrice int, err error)
}

type Repository struct {
	Order
	Store
	Payment
	Book
}

func NewRepository(db *gorm.DB, storeApiUri string, paymentApiUri string, bookApiUri string) *Repository {
	return &Repository{
		Order:   NewOrderPostgres(db),
		Store:   NewStoreApi(storeApiUri),
		Payment: NewPaymentApi(paymentApiUri),
		Book:    NewBookApi(bookApiUri),
	}
}
