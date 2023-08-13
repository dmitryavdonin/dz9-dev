package payment

import (
	"context"
	"order/internal/model"
)

type PaymentApi interface {
	DoPayment(ctx context.Context, info PaymentInfo) (result model.StatusResponse, err error)
}
