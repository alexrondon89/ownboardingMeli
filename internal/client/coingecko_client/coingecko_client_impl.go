package coingecko_client

import (
	"encoding/json"
	"ownboardingMeli/internal/client"
	"ownboardingMeli/internal/client/coingecko_client/dto"
	dto2 "ownboardingMeli/internal/client/dto"
	"ownboardingMeli/pkg/http"
	"strings"
)

type CoinGeckoClient struct {
	BaseUrl	string
}

func NewCoinGeckoClient(baseUrl string) *CoinGeckoClient{
	return &CoinGeckoClient{BaseUrl: baseUrl}
}

func (c *CoinGeckoClient) GetPriceFromClient(coin string, currency string) (*dto2.ClientResponse, error){
	endPoint := c.BaseUrl + coin
	httpResponse, err := http.GetRequest(endPoint)

	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != 200{
		return nil, client.ErrorBadRequest
	}

	coinGeckoResponse := dto.NewCoinGeckoResponse()
	clientResponse := dto2.NewClientResponse()
	err = json.Unmarshal(httpResponse.Body, coinGeckoResponse)
	if err != nil {
		return nil, err
	}

	if _, ok:= coinGeckoResponse.MarketData.CurrentPrice[currency]; !ok{
		return nil, client.ErrorCurrencyNotFound
	}

	clientResponse.Id = coinGeckoResponse.Id
	clientResponse.Price = coinGeckoResponse.MarketData.CurrentPrice[strings.ToLower(currency)]
	clientResponse.Currency = strings.ToUpper(currency)
	return clientResponse, nil
}