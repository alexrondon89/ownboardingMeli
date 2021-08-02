package service

import (
	service "ownboardingMeli/internal/service/dto"
)

type CryptoService interface {
	GetPrice(id string, currency string) (*service.CryptoResponse, error)
	GetListPrice(coins []string, currency string) ([]service.CryptoResponse, error)
}