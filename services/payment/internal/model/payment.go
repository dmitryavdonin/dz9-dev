package model

import (
	"time"
)

type Payment struct {
	ID         int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	OrderId    int       `gorm:"type:integer;not null" json:"order_id,omitempty"`
	Money      int       `gorm:"type:integer" json:"money"`
	UserId     int       `gorm:"type:integer;not null" json:"user_id,omitempty"`
	Status     string    `json:"status"`
	Reason     string    `json:"reason"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt time.Time `gorm:"not null" json:"modified_at,omitempty"`
}

type NewPayment struct {
	UserId  int `json:"user_id,omitempty"`
	OrderId int `json:"order_id,omitempty"`
	Money   int `json:"money"`
}

type CancelPayment struct {
	OrderId int    `json:"order_id,omitempty"`
	Reason  string `json:"reason"`
}
