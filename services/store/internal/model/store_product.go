package model

import (
	"time"
)

type Product struct {
	ID         int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	InStock    int       `gorm:"type:integer;not null" json:"in_stock,omitempty"`
	Price      int       `gorm:"type:integer;not null" json:"price,omitempty"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt time.Time `gorm:"not null" json:"modified_at,omitempty"`
}
