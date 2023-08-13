package repository

import (
	"context"
	"order/internal/repository/book/dto"
	"strconv"
)

type BookApi struct {
	baseUrl string
}

func NewBookApi(baseUrl string) *BookApi {
	return &BookApi{baseUrl: baseUrl}
}

func (r *BookApi) GetBookPrice(ctx context.Context, bookId int) (pookPrice int, err error) {

	response := &dto.GetBookPriceResponse{}

	_, err = sendGetRequest(r.baseUrl+"/"+strconv.Itoa(bookId), "application/json", response)
	if err != nil {
		return 0, err
	}
	return response.Price, nil
}
