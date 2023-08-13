package model

import (
	"time"
)

type StoreOrder struct {
	ID         int       `gorm:"type:integer;primary_key" json:"id,omitempty"`
	OrderId    int       `gorm:"type:integer;not null" json:"order_id"`
	BookId     int       `gorm:"type:integer;not null" json:"book_id"`
	Quantity   int       `gorm:"type:integer;not null" json:"quantity"`
	Status     string    `json:"status"`
	Reason     string    `json:"reason"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at,omitempty"`
	ModifiedAt time.Time `gorm:"not null" json:"modified_at,omitempty"`
}

type CancelOrder struct {
	OrderId int    `json:"order_id"`
	Reason  string `json:"reason"`
}
