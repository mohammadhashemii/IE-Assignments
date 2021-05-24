package model

import "time"

var Wallets []Wallet

type Wallet struct {
	Name        string    `json:"wname"`
	Balance     float64   `json:"balance"`
	Coins       []Coin    `json:"coins"`
	LastUpdated time.Time `json:"last_updated"`
}