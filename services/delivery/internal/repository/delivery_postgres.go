package repository

import (
	"delivery/internal/model"
	"time"

	"gorm.io/gorm"
)

type DelveryPostgres struct {
	db *gorm.DB
}

func NewDeliveryPostgres(db *gorm.DB) *DelveryPostgres {
	return &DelveryPostgres{db: db}
}

func (r *DelveryPostgres) Create(pay model.Delivery) (int, error) {
	result := r.db.Create(&pay)
	if result.Error != nil {
		return 0, result.Error
	}

	return pay.ID, nil
}

func (r *DelveryPostgres) GetById(orderId int) (model.Delivery, error) {
	var order model.Delivery
	result := r.db.First(&order, "order_id = ?", orderId)
	return order, result.Error
}

func (r *DelveryPostgres) GetAll(limit int, offset int) ([]model.Delivery, error) {
	var items []model.Delivery
	result := r.db.Limit(limit).Offset(offset).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, result.Error
}

func (r *DelveryPostgres) Delete(orderId int) error {
	result := r.db.Delete(&model.Delivery{}, "order_id = ?", orderId)
	return result.Error
}

func (r *DelveryPostgres) Update(input model.Delivery) error {
	var updatedItem model.Delivery
	result := r.db.First(&updatedItem, "order_id = ?", input.OrderId)
	if result.Error != nil {
		return result.Error
	}
	now := time.Now()
	itemToUpdate := model.Delivery{
		Status:     input.Status,
		Reason:     input.Reason,
		ModifiedAt: now,
	}
	result = r.db.Model(&updatedItem).Select("status", "reason", "modified_at").Updates(itemToUpdate)
	return result.Error
}
