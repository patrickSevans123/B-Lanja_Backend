package models

type Transaction struct {
	UserID uint    `json:"user_id"`
	Amount float64 `json:"amount"`
}
