package api

import (
	"ownboardingMeli/internal/api/dto"
)

type CryptoService interface {
	GetPrice(id string, currency string) (*dto.CryptoResponse, error)
}
