package repository

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"payment/internal/model"
	"payment/internal/repository/user/dto"
)

type UserApi struct {
	baseUrl string
}

func NewUserApi(baseUrl string) *UserApi {
	return &UserApi{baseUrl: baseUrl}
}

func (r *UserApi) GetBalance(ctx context.Context, user_id int) (ub model.UserBalance, err error) {
	request := &dto.CreateGetBalanceRequest{
		UserId: user_id,
	}

	response := &dto.CreateGetBalanceResponse{}

	_, err = sendRequest(r.baseUrl+"/user", http.MethodGet, "application/json", request, response)
	if err != nil {

		return model.UserBalance{
			Status: "failed",
			Reason: err.Error(),
		}, err
	}

	return model.UserBalance{
		UserId:  response.UserId,
		Balance: response.Balance,
	}, nil
}

func (r *UserApi) UpdateBalance(ctx context.Context, ub model.UserBalance) error {
	request := &dto.CreateUpdateBalanceRequest{
		UserId:  ub.UserId,
		Balance: ub.Balance,
	}

	response := &dto.CreateUpdateBalanceResponse{}

	_, err := sendRequest(r.baseUrl+"/user", http.MethodPut, "application/json", request, response)

	return err
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
