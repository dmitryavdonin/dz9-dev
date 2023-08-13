package user

import (
	"context"
)

type UserApi interface {
	GetBalance(ctx context.Context, user_id int) (int, error)
	UpdateBalance(ctx context.Context, user_id int, balance int) error
}
