package model

import (
	"time"
)

type StoreProduct struct {
	ID         int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	ProductId  int       `gorm:"type:integer;not null" json:"order_id,omitempty"`
	InStock    int       `gorm:"type:integer;not null" json:"in_stock,omitempty"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt time.Time `gorm:"not null" json:"modified_at,omitempty"`
}
