package dto

import (
	client "ownboardingMeli/internal/client/coingecko_client/dto"
	"strings"
)

type CryptoResponse struct {
	Id        	string      `json:"id"`
	Content		*Content	`json:"content,omitempty"`
	Partial 	bool		`json:"partial"`
}

type Content struct {
	Price		float64	`json:"price"`
	Currency	string		`json:"currency"`
}

type ListCryptoResponse struct {
	Items []CryptoResponse		`json:"items"`
}

func BuildCryptoResponse(clientResponse *client.CoinGeckoResponse, currency string) *CryptoResponse{
	return &CryptoResponse{
			Id: clientResponse.Id,
			Content: &Content{
				Price: clientResponse.MarketData.CurrentPrice[strings.ToLower(currency)],
				Currency: strings.ToUpper(currency),
			},
			Partial: false,
		}
}

func BuildPartialResponse(id string) *CryptoResponse{
	return &CryptoResponse{
		Id: id,
		Partial: true,
	}
}


