package model

import (
	"time"
)

type Book struct {
	ID         int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	Title      string    `gorm:"not null" json:"title,omitempty"`
	Author     string    `gorm:"not null" json:"author,omitempty"`
	Price      int       `gorm:"type:integer;not null" json:"price,omitempty"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt time.Time `gorm:"not null" json:"modified_at,omitempty"`
}
