package service

import (
	"context"
	"payment/internal/service/adapters/user"
)

type UserService struct {
	repo user.UserApi
}

func NewUserService(repo user.UserApi) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetBalance(ctx context.Context, user_id int) (int, error) {
	return s.repo.GetBalance(ctx, user_id)
}

func (s *UserService) UpdateBalance(ctx context.Context, user_id int, balance int) error {
	return s.repo.UpdateBalance(ctx, user_id, balance)
}
