package model

import (
	"time"
)

type Order struct {
	ID              int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	ProductId       int       `gorm:"type:integer;not null" json:"product_id,omitempty"`
	Quantity        int       `gorm:"type:integer;not null" json:"quantity,omitempty"`
	Price           int       `gorm:"type:integer;not null" json:"price,omitempty"`
	DeliveryAddress string    `gorm:"not null" json:"delivery_address,omitempty"`
	DeliveryDate    time.Time `gorm:"not null" json:"delivery_date,omitempty"`
	Status          string    `gorm:"not null" json:"status,omitempty"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt      time.Time `gorm:"not null" json:"modified_at,omitempty"`
}

type NewOrder struct {
	ProductId       int    `json:"product_id,omitempty"`
	Quantity        int    `json:"quantity,omitempty"`
	Price           int    `json:"price,omitempty"`
	DeliveryAddress string `json:"delivery_address,omitempty"`
	DeliveryDate    string `json:"delivery_date,omitempty"`
}

type UpdateOrder struct {
	ProductId       int       `json:"product_id,omitempty"`
	Quantity        int       `json:"quantity,omitempty"`
	Price           int       `json:"price,omitempty"`
	DeliveryAddress string    `json:"delivery_address,omitempty"`
	DeliveryDate    time.Time `json:"delivery_date,omitempty"`
	Status          string    `json:"status,omitempty"`
	CreatedAt       time.Time `json:"created_at,omitempty"`
	ModifiedAt      time.Time `json:"modified_at,omitempty"`
}
