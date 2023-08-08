package dto

type CreateStoreOrderRequest struct {
	OrderId  int `json:"order_id"`
	BookId   int `json:"book_id"`
	Quantity int `json:"quantity"`
}

type CreateStoreOrderResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}
