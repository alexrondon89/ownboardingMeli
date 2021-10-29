package http

import (
	"encoding/json"
	"ownboardingMeli/internal/client"
	"ownboardingMeli/pkg/http/service"
	"strings"
)

type CoinGeckoClient struct {
	BaseUrl	string
}

type CoinGeckoResponse struct {
	Id         string       `json:"id"`
	MarketData CurrentPrice `json:"market_data"`
}

type CurrentPrice struct {
	CurrentPrice	map[string]float64  `json:"current_price"`
}

func NewCoinGeckoResponse() *CoinGeckoResponse {
	return &CoinGeckoResponse{}
}

func NewCoinGeckoClient(baseUrl string) *CoinGeckoClient {
	return &CoinGeckoClient{BaseUrl: baseUrl}
}

func (c *CoinGeckoClient) GetPriceFromClient(coin string, currency string) (*client.Response, error){
	endPoint := c.BaseUrl + coin
	httpResponse, err := service.GetRequest(endPoint)

	if err != nil {
		return nil, err
	}

	if ok := service.CheckStatusCode200(httpResponse.StatusCode); !ok{
		return nil, client.ErrorBadRequest
	}

	coinGeckoResponse := NewCoinGeckoResponse()
	err = json.Unmarshal(httpResponse.Body, coinGeckoResponse)
	if err != nil {
		return nil, err
	}

	if _, ok:= coinGeckoResponse.MarketData.CurrentPrice[currency]; !ok{
		return nil, client.ErrorCurrencyNotFound
	}

	clientResponse := client.NewClientResponse()
	c.GenerateClientResponse(clientResponse, coinGeckoResponse, currency)
	return clientResponse, nil
}

func (c *CoinGeckoClient) GenerateClientResponse (resp *client.Response, geckResp *CoinGeckoResponse, currency string) {
	resp.Id = geckResp.Id
	resp.Price = geckResp.MarketData.CurrentPrice[strings.ToLower(currency)]
	resp.Currency = strings.ToUpper(currency)
}