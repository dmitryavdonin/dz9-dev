package repository

import (
	"store/internal/model"
	"time"

	"gorm.io/gorm"
)

type StoreOrderPostgres struct {
	db *gorm.DB
}

func NewStoreOrderPostgres(db *gorm.DB) *StoreOrderPostgres {
	return &StoreOrderPostgres{db: db}
}

func (r *StoreOrderPostgres) Create(order model.StoreOrder) (int, error) {
	result := r.db.Create(&order)
	if result.Error != nil {
		return 0, result.Error
	}
	return order.ID, nil
}

func (r *StoreOrderPostgres) GetById(order_id int) (model.StoreOrder, error) {
	var order model.StoreOrder
	result := r.db.First(&order, "order_id = ?", order_id)
	return order, result.Error
}

func (r *StoreOrderPostgres) GetAll(limit int, offset int) ([]model.StoreOrder, error) {
	var items []model.StoreOrder
	result := r.db.Limit(limit).Offset(offset).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, result.Error
}

func (r *StoreOrderPostgres) Delete(order_id int) error {
	result := r.db.Delete(&model.StoreOrder{}, "order_id = ?", order_id)
	return result.Error
}

func (r *StoreOrderPostgres) Update(order_id int, input model.StoreOrder) error {
	var updatedItem model.StoreOrder
	result := r.db.First(&updatedItem, "order_id = ?", order_id)
	if result.Error != nil {
		return result.Error
	}
	now := time.Now()
	itemToUpdate := model.StoreOrder{
		Status:     input.Status,
		Reason:     input.Reason,
		ModifiedAt: now,
	}
	result = r.db.Model(&updatedItem).Updates(itemToUpdate)
	return result.Error
}

func (r *StoreOrderPostgres) AlreadyExists(order_id int) bool {
	var order model.StoreOrder
	result := r.db.First(&order, "order_id = ?", order_id)
	return result.Error == nil
}
