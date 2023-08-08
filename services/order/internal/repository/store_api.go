package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"order/internal/model"
	"order/internal/repository/store/dto"
	"order/internal/service/adapters/store"
)

type StoreApi struct {
	baseUrl string
}

func NewStoreApi(baseUrl string) *StoreApi {
	return &StoreApi{baseUrl: baseUrl}
}

func (r *StoreApi) CreateStoreOrder(ctx context.Context, storeOrder store.StoreOrderInfo) (statusResponse *model.StatusResponse, err error) {
	request := &dto.CreateStoreOrderRequest{
		OrderId:  storeOrder.OrderId,
		BookId:   storeOrder.BookId,
		Quantity: storeOrder.Quantity,
	}

	response := &dto.CreateStoreOrderResponse{}

	_, err = sendRequest(r.baseUrl+"/order", http.MethodPost, "application/json", request, response)
	if err != nil {
		return
	}

	return &model.StatusResponse{
		Status: response.Status,
		Reason: response.Reason,
	}, nil
}

func sendRequest(url, method, contentType string, data interface{}, response interface{}) (code int, err error) {
	bodyBuffer, err := prepareBody(data)
	if err != nil {
		return
	}
	req, err := http.NewRequest(method, url, bodyBuffer)
	if err != nil {
		return 0, err
	}

	req.Header.Set("Content-Type", contentType)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, fmt.Errorf("no body in response")
	}

	if resp.StatusCode > 201 {
		return resp.StatusCode, fmt.Errorf("bad response code from server %s - code %d, body of response: %s", url, code, string(body))
	}

	if response != nil {
		err = json.Unmarshal(body, response)
		if err != nil {
			return 0, fmt.Errorf("error decode response from %s - error %s, body of response: %s", url, err.Error(), string(body))
		}
	}

	return resp.StatusCode, err
}

func prepareBody(data interface{}) (buffer *bytes.Buffer, err error) {
	var sendData []byte
	if data != nil {
		sendData, err = json.Marshal(data)
		if err != nil {
			return
		}
		buffer = bytes.NewBuffer(sendData)
	}
	return
}
