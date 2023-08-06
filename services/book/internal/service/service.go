package service

import (
	"book/internal/model"
	"book/internal/repository"
)

type Book interface {
	Create(input model.Book) (int, error)
	GetById(id int) (model.Book, error)
	GetAll(limit int, offset int) ([]model.Book, error)
	Delete(id int) error
	Update(id int, input model.Book) error
}

type Service struct {
	Book
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Book: NewBookService(repos.Book),
	}
}
