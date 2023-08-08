package store

import (
	"context"
	"order/internal/model"
)

type StoreApi interface {
	CreateStoreOrder(ctx context.Context, storeOrderInfo StoreOrderInfo) (statusRespone *model.StatusResponse, err error)
}
