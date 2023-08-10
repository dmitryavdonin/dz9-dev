package repository

import (
	"payment/internal/model"
	"time"

	"gorm.io/gorm"
)

type PaymentPostgres struct {
	db *gorm.DB
}

func NewPaymentPostgres(db *gorm.DB) *PaymentPostgres {
	return &PaymentPostgres{db: db}
}

func (r *PaymentPostgres) Create(pay model.Payment) (int, error) {
	result := r.db.Create(&pay)
	if result.Error != nil {
		return 0, result.Error
	}

	return pay.ID, nil
}

func (r *PaymentPostgres) GetById(orderId int) (model.Payment, error) {
	var order model.Payment
	result := r.db.First(&order, "id = ?", orderId)
	return order, result.Error
}

func (r *PaymentPostgres) GetAll(limit int, offset int) ([]model.Payment, error) {
	var items []model.Payment
	result := r.db.Limit(limit).Offset(offset).Find(&items)
	if result.Error != nil {
		return nil, result.Error
	}
	return items, result.Error
}

func (r *PaymentPostgres) Delete(orderId int) error {
	result := r.db.Delete(&model.Payment{}, "id = ?", orderId)
	return result.Error
}

func (r *PaymentPostgres) Update(input model.Payment) error {
	var updatedItem model.Payment
	result := r.db.First(&updatedItem, "id = ?", input.OrderId)
	if result.Error != nil {
		return result.Error
	}
	now := time.Now()
	itemToUpdate := model.Payment{
		Status:     input.Status,
		Reason:     input.Reason,
		ModifiedAt: now,
	}
	result = r.db.Model(&updatedItem).Select("status", "reason", "modified_at").Updates(itemToUpdate)
	return result.Error
}
