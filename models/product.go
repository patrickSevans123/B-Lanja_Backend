package models

type Product struct {
	ID 	   uint    `json:"id"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Stock  int     `json:"stock"`
	UserID uint    `json:"user_id"`
}
