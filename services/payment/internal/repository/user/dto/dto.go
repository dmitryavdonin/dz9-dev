package dto

type CreateGetBalanceRequest struct {
	UserId int `json:"user_id"`
}

type CreateGetBalanceResponse struct {
	UserId  int    `json:"user_id"`
	Balance int    `json:"balance"`
	Status  string `json:"status"`
	Reason  string `json:"reason"`
}

type CreateUpdateBalanceRequest struct {
	UserId  int `json:"user_id"`
	Balance int `json:"balance"`
}

type CreateUpdateBalanceResponse struct {
	UserId  int    `json:"user_id"`
	Balance int    `json:"balance"`
	Status  string `json:"status"`
	Reason  string `json:"reason"`
}
