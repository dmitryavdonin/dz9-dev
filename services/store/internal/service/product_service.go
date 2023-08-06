package service

import (
	"store/internal/model"
	"store/internal/repository"
)

type ProductService struct {
	repo repository.Product
}

func NewProductService(repo repository.Product) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) Create(input model.Product) (int, error) {
	return s.repo.Create(input)
}

func (s *ProductService) GetById(id int) (model.Product, error) {
	return s.repo.GetById(id)
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *ProductService) Update(id int, input model.Product) error {
	return s.repo.Update(id, input)
}
