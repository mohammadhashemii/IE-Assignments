package model

type Coin struct {
	Name   string  `json:"name"`
	Symbol string  `json:"symbol"`
	Amount float64 `json:"amount"`
	Rate   float64 `json:"rate"`
}