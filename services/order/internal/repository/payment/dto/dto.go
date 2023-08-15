package dto

type DoPaymentRequest struct {
	OrderId int `json:"order_id"`
	UserId  int `json:"user_id"`
	Money   int `json:"money"`
}

type DoPaymentResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type CancelPaymentRequest struct {
	OrderId int    `json:"order_id"`
	Reason  string `json:"reason"`
}

type CancelPaymentResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}
