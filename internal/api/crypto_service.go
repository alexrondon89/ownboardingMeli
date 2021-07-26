package api

import (
	"ownboardingMeli/internal/api/coingecko_service/dto"
	service "ownboardingMeli/internal/api/dto"

)

type CryptoService interface {
	GetPrice(id string, currency string) (*service.CryptoResponse, error)
	GetListPrice(coins []string, currency string) ([]*dto.ChannelInfo, error)
}
