package model

type UserBalance struct {
	UserId  int    `json:"user_id"`
	Balance int    `json:"balance"`
	Status  string `json:"status"`
	Reason  string `json:"reason"`
}
