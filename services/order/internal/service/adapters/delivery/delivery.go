package delivery

import (
	"context"
	"order/internal/model"
)

type DeliveryApi interface {
	DoDelivery(ctx context.Context, info DeliveryInfo) (result model.StatusResponse, err error)
}
