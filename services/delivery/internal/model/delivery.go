package model

import (
	"time"
)

type Delivery struct {
	ID              int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	OrderId         int       `gorm:"type:integer;not null" json:"order_id,omitempty"`
	UserId          int       `gorm:"type:integer;not null" json:"user_id,omitempty"`
	DeliveryAddress string    `json:"delivery_address"`
	DeliveryDate    time.Time `json:"delivery_date"`
	Status          string    `json:"status"`
	Reason          string    `json:"reason"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt      time.Time `gorm:"not null" json:"modified_at,omitempty"`
}

type NewDelivery struct {
	UserId          int    `json:"user_id,omitempty"`
	OrderId         int    `json:"order_id,omitempty"`
	DeliveryAddress string `json:"delivery_address"`
	DeliveryDate    string `json:"delivery_date"`
}

type CancelDelivery struct {
	OrderId int    `json:"order_id,omitempty"`
	Reason  string `json:"reason"`
}
