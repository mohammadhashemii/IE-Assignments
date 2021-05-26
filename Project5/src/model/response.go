package model


type Wallet_POST_PUT_DELETE_response struct {
	Name        string    `json:"name"`
	Balance     float64   `json:"balance"`
	Coins       []Coin    `json:"coins"`
	LastUpdated string `json:"last_updated"`
	Code		int		  `json:"code"`
	Message		string	  `json:"message"`
}

type Wallet_GET_response struct {
	Size        int    		`json:"size"`
	Wallets     []Wallet   	`json:"wallets"`
	Code		int		  	`json:"code"`
	Message		string	  	`json:"message"`
}

type Coin_POST_PUT_DELETE_response struct {
	Name        string    	`json:"name"`
	Symbol     	string   	`json:"symbol"`
	Amount      float64    	`json:"amount"`
	Rate		float64 	`json:"rate"`
	Code		int		  	`json:"code"`
	Message		string	  	`json:"message"`
}

type Wallet_GET_by_name_response struct {
	Name        string    `json:"name"`
	Balance     float64   `json:"balance"`
	Coins       []Coin    `json:"coins"`
	LastUpdated string `json:"last_updated"`
	Code		int		  `json:"code"`
	Message		string	  `json:"message"`
}

type Error_response struct {
	Code 		int 		`json:"code"`
	Message		string		`json:"message"`
}
