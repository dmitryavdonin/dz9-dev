package repository

import (
	"store/internal/model"

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

func (r *StoreOrderPostgres) GetByOrderId(id int) (model.StoreOrder, error) {
	var order model.StoreOrder
	result := r.db.First(&order, "order_id = ?", id)
	return order, result.Error
}

func (r *StoreOrderPostgres) DeleteByOrderId(order_id int) error {
	result := r.db.Delete(&model.StoreOrder{}, "order_id = ?", order_id)
	return result.Error
}
