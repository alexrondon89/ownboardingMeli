package client

import (
	"errors"
)

type PriceClient interface {
	GetPriceFromClient(coin string, currency string) (*Response, error)
}

// ErrorBadRequest appear when status code is not 200
var ErrorBadRequest = errors.New("status code different to 200")

// ErrorCurrencyNotFound appear whe currency is not present in response
var ErrorCurrencyNotFound = errors.New("Currency not found in client response ")

type Response struct {
	Id			string
	Price 		float64
	Currency 	string
}

func NewClientResponse() *Response {
	return &Response{}
}