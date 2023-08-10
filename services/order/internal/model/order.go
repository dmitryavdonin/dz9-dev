package model

import (
	"time"
)

type Order struct {
	ID              int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	BookId          int       `gorm:"type:integer;not null" json:"book_id,omitempty"`
	Quantity        int       `gorm:"type:integer;not null" json:"quantity,omitempty"`
	UserId          int       `gorm:"type:integer;not null" json:"user_id,omitempty"`
	DeliveryAddress string    `json:"delivery_address"`
	DeliveryDate    time.Time `json:"delivery_date"`
	Status          string    `json:"status"`
	Reason          string    `json:"reason"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt      time.Time `gorm:"not null" json:"modified_at,omitempty"`
}

type NewOrder struct {
	UserId          int    `json:"user_id,omitempty"`
	BookId          int    `json:"book_id,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
	DeliveryAddress string `json:"delivery_address"`
	DeliveryDate    string `json:"delivery_date"`
}
