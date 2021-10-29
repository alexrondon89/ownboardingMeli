package coingecko

import (
	"log"
	"ownboardingMeli/internal/client"
	"ownboardingMeli/internal/platform/builder"
	"ownboardingMeli/internal/service"
	"sync"
)

type ServiceCoinGecko struct {
	CoinGeckoClient client.PriceClient
}

func NewCoinGeckoService(client client.PriceClient) *ServiceCoinGecko {
	return &ServiceCoinGecko{CoinGeckoClient: client }
}

func (s *ServiceCoinGecko) GetPrice(coin string, currency string) (*service.CryptoResponse, error){
	clientResponse , err := s.CoinGeckoClient.GetPriceFromClient(coin, currency)
	if err != nil {
		return builder.PartialResponse(coin), err
	}

	response := builder.SuccessfulResponse(clientResponse, currency)
	return response, nil
}

func (s *ServiceCoinGecko) GetListPrice(coins []string, currency string) ([]service.CryptoResponse, error){
	n := len(coins)
	var ListResponse []service.CryptoResponse
	channel:= make(chan service.CryptoResponse, n) //concurrency
	wg:= sync.WaitGroup{}
	wg.Add(n)

	for _, v := range coins{
		go s.ReadCoinPrices(v, currency, &wg, channel)
	}

	wg.Wait() // wait for all concurrences
	close(channel)

	for c:= range channel{
		ListResponse = append(ListResponse, c)
	}

	return ListResponse, nil
}

func (s *ServiceCoinGecko) ReadCoinPrices(coin string, currency string, wg *sync.WaitGroup, c chan <- service.CryptoResponse){
	defer wg.Done()
	defer s.Recover(coin, c)

	clientResponse , err := s.CoinGeckoClient.GetPriceFromClient(coin, currency)

	// Only use to test Recover method after a panic
	if coin == "bitcoin"{
		panic("bitcoin error")
	}

	if err != nil {
		c<- *builder.PartialResponse(coin)
		return
	}

	c <- *builder.SuccessfulResponse(clientResponse, currency)
}

func (s *ServiceCoinGecko) Recover(coin string, c chan <- service.CryptoResponse){
	if r := recover(); r!= nil{
		log.Println("panic occurred: ",r)
		c <- *builder.PartialResponse(coin)
	}
}
