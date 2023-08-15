package repository

import (
	"context"
	"order/internal/model"
	"order/internal/service/adapters/delivery"
	"order/internal/service/adapters/payment"
	"order/internal/service/adapters/store"

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
	CancelPayment(ctx context.Context, orderId int, reason string) error
}

type Book interface {
	GetBookPrice(ctx context.Context, bookId int) (bookPrice int, err error)
}

type Delivery interface {
	DoDelivery(ctx context.Context, info delivery.DeliveryInfo) (statusRespone model.StatusResponse, err error)
}

type Repository struct {
	Order
	Store
	Payment
	Book
	Delivery
}

func NewRepository(db *gorm.DB, storeApiUri string, paymentApiUri string, bookApiUri string, deliveryApiUri string) *Repository {
	return &Repository{
		Order:    NewOrderPostgres(db),
		Store:    NewStoreApi(storeApiUri),
		Payment:  NewPaymentApi(paymentApiUri),
		Book:     NewBookApi(bookApiUri),
		Delivery: NewDeliveryApi(deliveryApiUri),
	}
}
