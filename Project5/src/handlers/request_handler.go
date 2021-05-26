package handlers

import (
	"Project5/src/model"
	"Project5/src/utils"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"io/ioutil"
	"net/http"
)

func Create_wallet(c echo.Context) error {

	new_wallet := model.Wallet{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		return Read_request_body_handler(err, c)
	}

	err = json.Unmarshal(b, &new_wallet)
	if err != nil{
		return Unmarshaling_handler(err, c)
	}
	err = json.Unmarshal(b, &new_wallet.X)
	delete(new_wallet.X, "name")
	if (len(new_wallet.X) > 0) {
		return Unknown_field_handler(c)
	}

	existing_wallet_idx, found := utils.Find_wallet(model.Wallets, new_wallet.Name)
	if found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This wallet(%s) is already added!\n", model.Wallets[existing_wallet_idx].Name))
	}

	new_wallet.Balance = 0.0
	new_wallet.LastUpdated = utils.Current_time()
	new_wallet.Coins = []model.Coin{}

	// add the new wallet to the wallets
	model.Wallets = append(model.Wallets, new_wallet)

	resp := model.Wallet_POST_PUT_DELETE_response{Name: new_wallet.Name,
													Balance: new_wallet.Balance,
													Coins: new_wallet.Coins,
													LastUpdated: new_wallet.LastUpdated,
													Code: http.StatusOK,
													Message: "Food added successfully"}

	return c.JSONPretty(http.StatusOK, resp, "	")
}

func Get_all_wallets(c echo.Context) error {
	resp := model.Wallet_GET_response{Size: len(model.Wallets),
										Wallets: model.Wallets,
										Code: http.StatusOK,
										Message: "All wallets received successfully!"}

	return c.JSONPretty(http.StatusOK, resp, "	")
}

func Edit_wallet(c echo.Context) error {

	wallet_idx, found := utils.Find_wallet(model.Wallets, c.Param("wname"))
	if !found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This wallet(%s) does not exist!\n", c.Param("wname")))
	}

	temp_wallet := model.Wallet{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		return Read_request_body_handler(err, c)
	}
	err = json.Unmarshal(b, &temp_wallet)
	if err != nil{
		return Unmarshaling_handler(err, c)
	}

	(&model.Wallets[wallet_idx]).Name = temp_wallet.Name
	(&model.Wallets[wallet_idx]).LastUpdated = utils.Current_time()

	wallet := model.Wallets[wallet_idx]
	resp := model.Wallet_POST_PUT_DELETE_response{Name: wallet.Name,
													Balance: wallet.Balance,
													Coins: wallet.Coins,
													LastUpdated: wallet.LastUpdated,
													Code: http.StatusOK,
													Message: "Wallet name changed successfully!"}

	return c.JSONPretty(http.StatusOK, resp, "	")
}

func Delete_wallet(c echo.Context) error {
	wallet_idx, found := utils.Find_wallet(model.Wallets, c.Param("wname"))
	if !found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This wallet(%s) does not exist!\n", c.Param("wname")))
	}

	deleted_wallet := model.Wallets[wallet_idx]
	model.Wallets = append(model.Wallets[:wallet_idx], model.Wallets[wallet_idx+1:]...)

	resp := model.Wallet_POST_PUT_DELETE_response{Name: deleted_wallet.Name,
												Balance: deleted_wallet.Balance,
												Coins: deleted_wallet.Coins,
												LastUpdated: deleted_wallet.LastUpdated,
												Code: http.StatusOK,
												Message: "Wallet deleted (logged out) successfully!"}

	return c.JSONPretty(http.StatusOK, resp, "	")
}

func Create_coin(c echo.Context) error {
	new_coin := model.Coin{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		return Read_request_body_handler(err, c)
	}
	err = json.Unmarshal(b, &new_coin)
	if err != nil{
		return Unmarshaling_handler(err, c)
	}

	err = json.Unmarshal(b, &new_coin.X)
	delete(new_coin.X, "name")
	delete(new_coin.X, "symbol")
	delete(new_coin.X, "amount")
	delete(new_coin.X, "rate")
	if (len(new_coin.X) > 0) {
		return Unknown_field_handler(c)
	}


	wallet_idx, found := utils.Find_wallet(model.Wallets, c.Param("wname"))
	if !found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This wallet(%s) does not exist!\n", c.Param("wname")))
	}

	existing_coin_idx, coin_found := utils.Find_coin((&model.Wallets[wallet_idx]).Coins, new_coin.Name)
	if coin_found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This coin(%s) is already added to %s!\n", model.Wallets[wallet_idx].Coins[existing_coin_idx].Symbol, model.Wallets[wallet_idx].Name))
	}

	(&model.Wallets[wallet_idx]).LastUpdated = utils.Current_time()
	(&model.Wallets[wallet_idx]).Coins = append((&model.Wallets[wallet_idx]).Coins, new_coin)
	(&model.Wallets[wallet_idx]).Balance += new_coin.Amount * new_coin.Rate

	resp := model.Coin_POST_PUT_DELETE_response{Name: new_coin.Name,
												Symbol: new_coin.Symbol,
												Amount: new_coin.Amount,
												Rate: new_coin.Rate,
												Code: http.StatusOK,
												Message: "Coin added successfully!"}

	return c.JSONPretty(http.StatusOK, resp, "	")
}

func Get_wallet_info(c echo.Context) error {
	wallet_idx, found := utils.Find_wallet(model.Wallets, c.Param("wname"))
	if !found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This wallet(%s) does not exist!\n", c.Param("wname")))
	}

	wallet := model.Wallets[wallet_idx]
	resp := model.Wallet_GET_by_name_response{Name: wallet.Name,
													Balance: wallet.Balance,
													Coins: wallet.Coins,
													LastUpdated: wallet.LastUpdated,
													Code: http.StatusOK,
													Message: "All coins received successfully!"}

	return c.JSONPretty(http.StatusOK, resp, "	")

}

func Edit_coin(c echo.Context) error {

	new_coin := model.Coin{}
	defer c.Request().Body.Close()
	b, err := ioutil.ReadAll(c.Request().Body)
	if err != nil{
		return Read_request_body_handler(err, c)
	}
	err = json.Unmarshal(b, &new_coin)
	if err != nil{
		return Unmarshaling_handler(err, c)
	}

	err = json.Unmarshal(b, &new_coin.X)
	delete(new_coin.X, "name")
	delete(new_coin.X, "symbol")
	delete(new_coin.X, "amount")
	delete(new_coin.X, "rate")
	if (len(new_coin.X) > 0) {
		return Unknown_field_handler(c)
	}

	resp := model.Coin_POST_PUT_DELETE_response{Name: new_coin.Name,
												Symbol: new_coin.Symbol,
												Amount: new_coin.Amount,
												Rate: new_coin.Rate,
												Code: http.StatusOK,
												Message: "Coin updated successfully!"}

	wallet_idx, found := utils.Find_wallet(model.Wallets, c.Param("wname"))
	if !found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This wallet(%s) does not exist!\n", c.Param("wname")))
	}

	coin_idx, coin_found := utils.Find_coin((&model.Wallets[wallet_idx]).Coins, c.Param("symbol"))
	if !coin_found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This coin(%s) does not exist!\n", c.Param("symbol")))
	}


	(&model.Wallets[wallet_idx]).LastUpdated = utils.Current_time()
	(&model.Wallets[wallet_idx]).Balance -= model.Wallets[wallet_idx].Coins[coin_idx].Amount * model.Wallets[wallet_idx].Coins[coin_idx].Rate
	(&model.Wallets[wallet_idx]).Balance += new_coin.Amount * new_coin.Rate

	(&model.Wallets[wallet_idx]).Coins[coin_idx] = new_coin

	return c.JSONPretty(http.StatusOK, resp, "	")
}

func Delete_coin(c echo.Context) error {

	wallet_idx, found := utils.Find_wallet(model.Wallets, c.Param("wname"))
	if !found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This wallet(%s) does not exist!\n", c.Param("wname")))
	}

	coin_idx, coin_found := utils.Find_coin((&model.Wallets[wallet_idx]).Coins, c.Param("symbol"))
	if !coin_found {
		return c.String(http.StatusBadRequest, fmt.Sprintf("This coin(%s) does not exist!\n", c.Param("symbol")))
	}

	deleted_coin := (&model.Wallets[wallet_idx]).Coins[coin_idx]
	(&model.Wallets[wallet_idx]).LastUpdated = utils.Current_time()
	(&model.Wallets[wallet_idx]).Balance -= model.Wallets[wallet_idx].Coins[coin_idx].Amount * model.Wallets[wallet_idx].Coins[coin_idx].Rate
	(&model.Wallets[wallet_idx]).Coins = append((&model.Wallets[wallet_idx]).Coins[:coin_idx], (&model.Wallets[wallet_idx]).Coins[coin_idx+1:]...)

	resp := model.Coin_POST_PUT_DELETE_response{Name: deleted_coin.Name,
												Symbol: deleted_coin.Symbol,
												Amount: deleted_coin.Amount,
												Rate: deleted_coin.Rate,
												Code: http.StatusOK,
												Message: "Coin deleted successfully!"}

	return c.JSONPretty(http.StatusOK, resp, "	")
}