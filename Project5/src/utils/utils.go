package utils

import (
	"Project5/src/model"
	"fmt"
	"time"
)

// Current_time A function to return the current time in arbitrary format
func Current_time() string{
	t := time.Now()
	formatted_time := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	return formatted_time
}

// Find_wallet takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find_wallet(wallets []model.Wallet, val string) (int, bool) {
	for i, item := range wallets {
		if item.Name == val {
			return i, true
		}
	}
	return -1, false
}

// Find_coin takes a slice and looks for an element in it. If found it will
// return it's key, otherwise it will return -1 and a bool of false.
func Find_coin(coins []model.Coin, val string) (int, bool) {
	for i, item := range coins {
		if item.Name == val || item.Symbol == val{
			return i, true
		}
	}
	return -1, false
}