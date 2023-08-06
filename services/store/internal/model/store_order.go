package model

import (
	"time"
)

type StoreOrder struct {
	ID         int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	OrderId    int       `gorm:"type:integer;not null" json:"order_id,omitempty"`
	ProductId  int       `gorm:"type:integer;not null" json:"product_id,omitempty"`
	Quantity   int       `gorm:"type:integer;not null" json:"quantity,omitempty"`
	Status     string    `gorm:"not null" json:"status,omitempty"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt time.Time `gorm:"not null" json:"modified_at,omitempty"`
}
