package dto

type CoinGeckoResponse struct {
	Id        	string       `json:"id"`
	MarketData 	CurrentPrice `json:"market_data"`
}

type CurrentPrice struct {
	CurrentPrice	map[string]float64  `json:"current_price"`
}

func NewCoinGeckResponse() *CoinGeckoResponse {
	return &CoinGeckoResponse{}
}