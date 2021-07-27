package api

import (
	"errors"
	service "ownboardingMeli/internal/api/dto"
)

type CryptoService interface {
	GetPrice(id string, currency string) (*service.CryptoResponse, error)
	GetListPrice(coins []string, currency string) ([]service.CryptoResponse, error)
}

var ErrorCurrencyNotExist = errors.New("currency is not available")
