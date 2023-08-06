package service

import (
	"book/internal/model"
	"book/internal/repository"
)

type BookService struct {
	repo repository.Book
}

func NewBookService(repo repository.Book) *BookService {
	return &BookService{repo: repo}
}

func (s *BookService) Create(input model.Book) (int, error) {
	return s.repo.Create(input)
}

func (s *BookService) GetById(id int) (model.Book, error) {
	return s.repo.GetById(id)
}

func (s *BookService) GetAll(limit int, offset int) ([]model.Book, error) {
	return s.repo.GetAll(limit, offset)
}

func (s *BookService) Delete(id int) error {
	return s.repo.Delete(id)
}

func (s *BookService) Update(id int, input model.Book) error {
	return s.repo.Update(id, input)
}
