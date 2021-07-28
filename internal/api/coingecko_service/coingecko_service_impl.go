package coingecko_service

import (
	"log"
	"ownboardingMeli/internal/api"
	service "ownboardingMeli/internal/api/dto"
	"ownboardingMeli/internal/client/coingecko_client"
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
		return service.BuildPartialResponse(id), err
	}

	if _, ok := clientResponse.MarketData.CurrentPrice[currency]; !ok {
		return service.BuildPartialResponse(id), api.ErrorCurrencyNotExist
	}

	response := service.BuildCryptoResponse(clientResponse, currency)
	return response, nil
}

func (s *CoinGeckoService) GetListPrice(coins []string, currency string) ([]service.CryptoResponse, error){
	n := len(coins)
	var ListResponse []service.CryptoResponse
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
	if id == "bitcoin"{
		panic("error de bitcoin")
	}

	if err != nil {
		c<- *service.BuildPartialResponse(id)
	}

	c<- *service.BuildCryptoResponse(clientResponse, currency)
}

func (s *CoinGeckoService) Recover(id string, c chan <- service.CryptoResponse){
	if r := recover(); r!= nil{
		log.Println("panic occurred: ",r)
		c <- *service.BuildPartialResponse(id)
	}
}
