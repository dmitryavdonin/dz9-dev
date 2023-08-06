package repository

import (
	"book/internal/model"
	"time"

	"gorm.io/gorm"
)

type BookPostgres struct {
	db *gorm.DB
}

func NewBookPostgres(db *gorm.DB) *BookPostgres {
	return &BookPostgres{db: db}
}

func (r *BookPostgres) Create(input model.Book) (int, error) {
	result := r.db.Create(&input)
	if result.Error != nil {
		return 0, result.Error
	}
	return input.ID, nil
}

func (r *BookPostgres) GetById(id int) (model.Book, error) {
	var item model.Book
	result := r.db.First(&item, "id = ?", id)
	return item, result.Error
}

func (r *BookPostgres) GetAll(limit int, offset int) ([]model.Book, error) {
	var items []model.Book
	result := r.db.Limit(limit).Offset(offset).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, result.Error
}

func (r *BookPostgres) Delete(id int) error {
	result := r.db.Delete(&model.Book{}, "id = ?", id)
	return result.Error
}

func (r *BookPostgres) Update(id int, input model.Book) error {
	var updated model.Book
	result := r.db.First(&updated, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	now := time.Now()
	bookToUpdate := model.Book{
		Title:      input.Title,
		Author:     input.Author,
		Price:      input.Price,
		CreatedAt:  updated.CreatedAt,
		ModifiedAt: now,
	}
	result = r.db.Model(&updated).Updates(bookToUpdate)
	return result.Error
}
