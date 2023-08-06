package repository

import (
	"book/internal/model"

	"gorm.io/gorm"
)

type Book interface {
	Create(item model.Book) (int, error)
	GetById(id int) (model.Book, error)
	GetAll(limit int, offset int) ([]model.Book, error)
	Delete(id int) error
	Update(id int, item model.Book) error
}

type Repository struct {
	Book
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Book: NewBookPostgres(db),
	}
}
