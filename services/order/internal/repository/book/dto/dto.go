package dto

type GetBookPriceRequest struct {
	BookId int `json:"book_id"`
}

type GetBookPriceResponse struct {
	Price int `json:"price"`
}
