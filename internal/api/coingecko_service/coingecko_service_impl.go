package coingecko_service

import (
	"encoding/json"
	dto2 "ownboardingMeli/internal/api/coingecko_service/dto"
	"ownboardingMeli/internal/api/dto"
	"ownboardingMeli/pkg/http"
	"strings"
)

type CoinGeckoService struct {
	BaseUrl string
}

func NewCoinGeckoService(url string) *CoinGeckoService{
	return &CoinGeckoService{url}
}

func (s *CoinGeckoService) GetPrice(id string, currency string) (*dto.CryptoResponse, error){
	url := s.BaseUrl + id
	bodyByte, err := http.GetRequest(url)
	if err != nil {
		return nil, err
	}

	response, err := buildResponse(bodyByte, currency)
	return response, err
}

func buildResponse(bodyByte []byte, currency string) (*dto.CryptoResponse, error){
	var cryptoResponse dto.CryptoResponse
	var out dto2.CoinGeckoResponse
	err := json.Unmarshal(bodyByte, &out)

	if err != nil {
		return nil, err
	}

	cryptoResponse.Id = out.Id
	cryptoResponse.Content.Price = out.MarketData.CurrentPrice[currency]
	cryptoResponse.Content.Currency = strings.ToUpper(currency)
	cryptoResponse.Partial = false

	return &cryptoResponse, nil
}