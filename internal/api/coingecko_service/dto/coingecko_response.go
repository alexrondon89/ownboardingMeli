package dto

import (
	"ownboardingMeli/internal/client/coingecko_client/dto"
)

type ChannelInfo struct {
	CryptoResponse 	*dto.CoinGeckoResponse
	Error 			error
	Coin			string
	Currency 		string
}

func NewChannelInfo(body *dto.CoinGeckoResponse, err error, coin string, currency string) *ChannelInfo{
	return &ChannelInfo{CryptoResponse: body, Error: err, Coin: coin, Currency: currency}
}

