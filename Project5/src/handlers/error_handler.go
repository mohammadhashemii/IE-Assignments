package handlers

import (
	"Project5/src/model"
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

func Read_request_body_handler(err error, c echo.Context) error{
	log.Printf("Failed reading the request body!: %s\n", err)
	return c.JSONPretty(http.StatusInternalServerError, model.Error_response{Code: http.StatusInternalServerError, Message: "Internal server error!"}, "	")
}

func Unmarshaling_handler(err error, c echo.Context) error{
	log.Printf("Failed unmarshaling!: %s\n", err)
	return c.JSONPretty(http.StatusBadRequest, model.Error_response{Code: http.StatusBadRequest, Message: "The request JSON format is not valid!"}, "	")
}

func Unknown_field_handler(c echo.Context) error{
	log.Printf("Unknown fields found while unmarshaling!\n")
	return c.JSONPretty(http.StatusBadRequest, model.Error_response{Code: http.StatusBadRequest, Message: "You entered unsupported fields in the input JSON!"}, "	")
}

func Bad_endpoint_handler(c echo.Context) error{
	log.Printf("Method not allowed error!\n")
	return c.JSONPretty(http.StatusNotImplemented, model.Error_response{Code: http.StatusNotImplemented, Message: "This endpoint has not been yet implemented!"}, "	")
}