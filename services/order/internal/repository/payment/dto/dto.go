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
