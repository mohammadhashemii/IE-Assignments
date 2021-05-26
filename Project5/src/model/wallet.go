package model

var Wallets []Wallet

type Wallet struct {
	Name        string    				`json:"name"`
	Balance     float64   				`json:"balance"`
	Coins       []Coin    				`json:"coins"`
	LastUpdated string 					`json:"last_updated"`
	X 			map[string]interface{} 	`json:"-"` // Rest of the fields should go here.
}