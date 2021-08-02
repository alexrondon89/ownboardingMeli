package client

import (
	"errors"
	"ownboardingMeli/internal/client/dto"
)

type ClientInterface interface {
	GetPriceFromClient(coin string, currency string) (*dto.ClientResponse, error)
}

// ErrorBadRequest appear when status code is not 200
var ErrorBadRequest = errors.New("status code different to 200")

// ErrorCurrencyNotFound appear whe currency is not present in response
var ErrorCurrencyNotFound = errors.New("Currency not found in client response ")
