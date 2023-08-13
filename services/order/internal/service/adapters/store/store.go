package store

import (
	"context"
	"order/internal/model"
)

type StoreApi interface {
	PlaceOrderInStore(ctx context.Context, storeOrderInfo StoreOrderInfo) (statusRespone model.StatusResponse, err error)
	CancelOrderInStore(ctx context.Context, orderId int, reason string) error
}
