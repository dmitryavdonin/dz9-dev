package model

import (
	"time"
)

type StoreBook struct {
	ID         int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	BookId     int       `gorm:"type:integer;not null" json:"book_id,omitempty"`
	InStock    int       `gorm:"type:integer;" json:"in_stock"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt time.Time `gorm:"not null" json:"modified_at,omitempty"`
}
