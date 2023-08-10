package handler

type StatusResponse struct {
	Id     int    `json:"id"`
	Status string `json:"status"`
	Reason string `json:"reason"`
}
