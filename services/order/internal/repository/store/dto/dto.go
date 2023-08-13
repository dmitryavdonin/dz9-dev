package dto

type PlaceOrderInStoreRequest struct {
	OrderId  int `json:"order_id"`
	BookId   int `json:"book_id"`
	Quantity int `json:"quantity"`
}

type PlaceOrderInStoreResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}

type CancelOrderInStoreRequest struct {
	OrderId int    `json:"order_id"`
	Reason  string `json:"reason"`
}

type CancelOrderInStoreResponse struct {
	Status string `json:"status"`
	Reason string `json:"reason"`
}
