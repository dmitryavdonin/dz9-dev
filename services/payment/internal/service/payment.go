package service

import (
	"payment/internal/model"
	"payment/internal/repository"
)

type PaymentService struct {
	repo repository.Payment
}

func NewPaymentService(repo repository.Payment) *PaymentService {
	return &PaymentService{repo: repo}
}

func (s *PaymentService) Create(pay model.Payment) (int, error) {
	return s.repo.Create(pay)
}

func (s *PaymentService) GetById(id int) (model.Payment, error) {
	return s.repo.GetById(id)
}

func (s *PaymentService) GetAll(limit int, offset int) ([]model.Payment, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *PaymentService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *PaymentService) Update(pay model.Payment) error {
	return s.repo.Update(pay)
}
