package model

type UserBalance struct {
	Balance int    `json:"balance"`
	Status  string `json:"status"`
	Reason  string `json:"reason"`
}
