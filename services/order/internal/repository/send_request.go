package repository

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func sendGetRequest(url, contentType string, response interface{}) (code int, err error) {
	req, err := http.NewRequest("GET", url, nil)
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
	} else {
		buffer = nil
	}
	return
}
