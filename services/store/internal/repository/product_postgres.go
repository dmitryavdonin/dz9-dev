package repository

import (
	"store/internal/model"
	"time"

	"gorm.io/gorm"
)

type ProductPostgres struct {
	db *gorm.DB
}

func NewProductPostgres(db *gorm.DB) *ProductPostgres {
	return &ProductPostgres{db: db}
}

func (r *ProductPostgres) Create(sp model.StoreProduct) (int, error) {
	result := r.db.Create(&sp)
	if result.Error != nil {
		return 0, result.Error
	}
	return sp.ID, nil
}

func (r *ProductPostgres) GetById(id int) (model.StoreProduct, error) {
	var product model.StoreProduct
	result := r.db.First(&product, "id = ?", id)
	return product, result.Error
}

func (r *ProductPostgres) GetByProductId(id int) (model.StoreProduct, error) {
	var product model.StoreProduct
	result := r.db.First(&product, "product_id = ?", id)
	return product, result.Error
}

func (r *ProductPostgres) Delete(id int) error {
	result := r.db.Delete(&model.StoreProduct{}, "id = ?", id)
	return result.Error
}

func (r *ProductPostgres) Update(id int, input model.StoreProduct) error {
	var updatedProduct model.StoreProduct
	result := r.db.First(&updatedProduct, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	now := time.Now()
	productToUpdate := model.StoreProduct{
		InStock:    input.InStock,
		ProductId:  input.ProductId,
		CreatedAt:  updatedProduct.CreatedAt,
		ModifiedAt: now,
	}
	result = r.db.Model(&updatedProduct).Updates(productToUpdate)
	return result.Error
}
