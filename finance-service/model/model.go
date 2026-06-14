package model

type TopUpRequest struct {
	UserID string `json:"user_id"`
	Amount int    `json:"amount"`
}
