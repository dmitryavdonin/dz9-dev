package user

import (
	"context"
	"payment/internal/model"
)

type UserApi interface {
	GetBalance(ctx context.Context, user_id int) (model.UserBalance, error)
	UpdateBalance(ctx context.Context, ub model.UserBalance) error
}
