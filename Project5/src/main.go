package main

import (
	"Project5/src/handlers"
	"fmt"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("The server started listening!")
	e:= echo.New()

	e.POST("/wallets", handlers.Create_wallet)
	e.GET("/wallets", handlers.Get_all_wallets)
	e.PUT("/wallets/:wname", handlers.Edit_wallet)
	e.DELETE("/wallets/:wname", handlers.Delete_wallet)

	e.POST("/:wname/coins", handlers.Create_coin)
	e.GET("/:wname", handlers.Get_wallet_info)
	e.PUT("/:wname/:symbol", handlers.Edit_coin)
	e.DELETE("/:wname/:symbol", handlers.Delete_coin)

	// Bad endpoints
	e.Any("/*", handlers.Bad_endpoint_handler)

	e.Logger.Fatal(e.Start(":1323"))
}