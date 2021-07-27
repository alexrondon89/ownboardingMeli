package coingecko_service

import (
	"log"
	"ownboardingMeli/internal/api"
	service "ownboardingMeli/internal/api/dto"
	"ownboardingMeli/internal/client/coingecko_client"
	client "ownboardingMeli/internal/client/coingecko_client/dto"
	"strings"
	"sync"
)

type CoinGeckoService struct {
	CoinGeckoClient coingecko_client.CoinGeckoClient
}

func NewCoinGeckoService(client *coingecko_client.CoinGeckoClient) *CoinGeckoService{
	return &CoinGeckoService{CoinGeckoClient: *client }
}

func (s *CoinGeckoService) GetPrice(id string, currency string) (*service.CryptoResponse, error){
	clientResponse , err := s.CoinGeckoClient.GetCoinPrice(id)
	if err != nil {
		return s.buildPartialResponse(id), err
	}

	if _, ok := clientResponse.MarketData.CurrentPrice[currency]; !ok {
		return s.buildPartialResponse(id), api.ErrorCurrencyNotExist
	}

	response := s.buildResponse(clientResponse, currency)
	return response, nil
}

func (s *CoinGeckoService) GetListPrice(coins []string, currency string) ([]service.CryptoResponse, error){
	n := len(coins)
	ListResponse:= make([]service.CryptoResponse, n)
	channel:= make(chan service.CryptoResponse, n)
	wg:= sync.WaitGroup{}
	wg.Add(n)

	for _, v := range coins{
		go s.ReadCoinPrice(v, currency, &wg, channel)
	}

	wg.Wait()
	close(channel)

	for c:= range channel{
		ListResponse = append(ListResponse, c)
	}

	return ListResponse, nil
}

func (s *CoinGeckoService) ReadCoinPrice(id string, currency string, wg *sync.WaitGroup, c chan <- service.CryptoResponse){
	defer wg.Done()
	defer s.Recover(id, c)

	clientResponse , err := s.CoinGeckoClient.GetCoinPrice(id)
	if err != nil {
		c<- *s.buildPartialResponse(id)
	}

	if id == "bitcoin"{
		log.Println("imprime error de bitcoin")
		panic("error de bitcoin")
	}

	c<- *s.buildResponse(clientResponse, currency)

}

func (s *CoinGeckoService) buildResponse(clientResponse *client.CoinGeckoResponse, currency string) *service.CryptoResponse {
	return &service.CryptoResponse{
		Id: clientResponse.Id,
		Content: &service.Content{
			Price: clientResponse.MarketData.CurrentPrice[strings.ToLower(currency)],
			Currency: strings.ToUpper(currency),
		},
		Partial: false,
	}
}

func (s *CoinGeckoService) buildPartialResponse(id string) *service.CryptoResponse{
	return &service.CryptoResponse{
		Id: id,
		Partial: true,
	}
}

func (s *CoinGeckoService) Recover(id string, c chan <- service.CryptoResponse) {
	if r := recover(); r!= nil{
		log.Println("asigna en recover")
		c<- *s.buildPartialResponse("prueba")
	}
}
