package coingecko_service

import (
	"log"
	"ownboardingMeli/internal/api/coingecko_service/dto"
	service "ownboardingMeli/internal/api/dto"
	"ownboardingMeli/internal/client/coingecko_client"
	client "ownboardingMeli/internal/client/coingecko_client/dto"
	"strings"
)

var ListResponse [] *dto.ChannelInfo

type CoinGeckoService struct {
	CoinGeckoClient coingecko_client.CoinGeckoClient
}

func NewCoinGeckoService(client *coingecko_client.CoinGeckoClient) *CoinGeckoService{
	return &CoinGeckoService{CoinGeckoClient: *client }
}

func (s *CoinGeckoService) GetPrice(id string, currency string) (*service.CryptoResponse, error){
	clientResponse , err := s.CoinGeckoClient.GetCoinPrice(id)
	if err != nil {
		return nil, err
	}

	response, err := buildResponse(clientResponse, currency)
	return response, err
}

func (s *CoinGeckoService) GetListPrice(coins []string, currency string) ([]*dto.ChannelInfo, error){

	channelA := make(chan *dto.ChannelInfo)
	channelB := make(chan *dto.ChannelInfo)
	channelC := make(chan *dto.ChannelInfo)

	go s.ReadCoinPrice(coins[0], currency, channelA)
	go s.ReadCoinPrice(coins[1], currency, channelB)
	go s.ReadCoinPrice(coins[2], currency, channelC)
	a := <- channelA
	b := <- channelB
	c := <- channelC

	log.Println(a)
	log.Println(b)
	log.Println(c)

	ListResponse = append(ListResponse, a, b, c)
	log.Println("luego de realizar las consultas")
	log.Println(ListResponse)
	return ListResponse, nil
}

func (s *CoinGeckoService) ReadCoinPrice(coin string, currency string, c chan <- *dto.ChannelInfo){
	response , err := s.CoinGeckoClient.GetCoinPrice(coin)
	channelInfo := dto.NewChannelInfo(response, err, coin, currency)
	c <- channelInfo
}

func (s *CoinGeckoService) ListenCoinPrice(channelA, channelB , channelC <- chan client.CoinGeckoResponse) {
	for {
		log.Println("entra en ListenCoinPrice")
		select {
			case a := <- channelA:
				log.Println("entra en channelA", a)
			case b := <- channelB:
				log.Println("entra en channelB", b)
			case c := <- channelC:
				log.Println("entra en channelC", c)
		}
		log.Println(ListResponse)
	}
}

func buildResponse(clientResponse *client.CoinGeckoResponse, currency string) (*service.CryptoResponse, error){
	var cryptoResponse service.CryptoResponse

	cryptoResponse.Id = clientResponse.Id
	cryptoResponse.Content.Price = clientResponse.MarketData.CurrentPrice[currency]
	cryptoResponse.Content.Currency = strings.ToUpper(currency)
	cryptoResponse.Partial = false

	return &cryptoResponse, nil
}