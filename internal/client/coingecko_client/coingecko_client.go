package coingecko_client

import (
	"encoding/json"
	"errors"
	"ownboardingMeli/internal/client/coingecko_client/dto"
	"ownboardingMeli/pkg/http"
)

type CoinGeckoClient struct {
	BaseUrl	string
}

func NewCoinGeckoClient(baseUrl string) *CoinGeckoClient{
	return &CoinGeckoClient{BaseUrl: baseUrl}
}

func (c *CoinGeckoClient) GetCoinPrice(path string) ( *dto.CoinGeckoResponse,error){
	endPoint := c.BaseUrl + path
	httpResponse, err := http.GetRequest(endPoint)

	if err != nil {
		return nil, err
	}

	if httpResponse.StatusCode != 200{
		return nil, errors.New("status code different to 200")
	}

	coinGeckoResponse := dto.NewCoinGeckResponse()
	err = json.Unmarshal(httpResponse.Body, &coinGeckoResponse)
	if err != nil {
		return nil, err
	}

	return coinGeckoResponse, nil
}