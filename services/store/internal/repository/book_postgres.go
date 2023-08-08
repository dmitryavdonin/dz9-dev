package repository

import (
	"store/internal/model"
	"time"

	"gorm.io/gorm"
)

type StoreBookPostgres struct {
	db *gorm.DB
}

func NewStoreBookPostgres(db *gorm.DB) *StoreBookPostgres {
	return &StoreBookPostgres{db: db}
}

func (r *StoreBookPostgres) Create(sp model.StoreBook) (int, error) {
	result := r.db.Create(&sp)
	if result.Error != nil {
		return 0, result.Error
	}
	return sp.ID, nil
}

func (r *StoreBookPostgres) GetById(book_id int) (model.StoreBook, error) {
	var item model.StoreBook
	result := r.db.First(&item, "book_id = ?", book_id)
	return item, result.Error
}

func (r *StoreBookPostgres) GetAll(limit int, offset int) ([]model.StoreBook, error) {
	var items []model.StoreBook
	result := r.db.Limit(limit).Offset(offset).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, result.Error
}

func (r *StoreBookPostgres) Delete(book_id int) error {
	result := r.db.Delete(&model.StoreBook{}, "book_id = ?", book_id)
	return result.Error
}

func (r *StoreBookPostgres) Update(book_id int, input model.StoreBook) error {
	var updatedItem model.StoreBook
	result := r.db.First(&updatedItem, "book_id = ?", book_id)
	if result.Error != nil {
		return result.Error
	}
	now := time.Now()
	itemToUpdate := model.StoreBook{
		InStock:    input.InStock,
		ModifiedAt: now,
	}
	result = r.db.Model(&updatedItem).Select("in_stock", "modified_at").Updates(itemToUpdate)
	return result.Error
}

func (r *StoreBookPostgres) AlreadyExists(book_id int) bool {
	var order model.StoreOrder
	result := r.db.First(&order, "book_id = ?", book_id)
	return result.Error == nil
}
