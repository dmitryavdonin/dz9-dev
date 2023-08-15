package delivery

import "time"

type DeliveryInfo struct {
	OrderId         int
	UserId          int
	DeliveryAddress string
	DeliveryDate    time.Time
}

type DeliveryResult struct {
	Status string
	Result string
}
