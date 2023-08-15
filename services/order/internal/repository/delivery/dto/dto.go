package dto

type DoDeliveryRequest struct {
	OrderId         int    `json:"order_id"`
	UserId          int    `json:"user_id"`
	DeliveryAddress string `json:"delivery_address"`
	DeliveryDate    string `json:"delivery_date"`
}

type DoDeliveryResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type CancelDeliveryRequest struct {
	OrderId int    `json:"order_id"`
	Reason  string `json:"reason"`
}

type CancelDeliveryResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}
