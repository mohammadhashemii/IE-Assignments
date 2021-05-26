package model

type Coin struct {
	Name   	string  				`json:"name"`
	Symbol 	string  				`json:"symbol"`
	Amount 	float64 				`json:"amount"`
	Rate   	float64 				`json:"rate"`
	X 	 	map[string]interface{} 	`json:"-"` // Rest of the fields should go here.
}