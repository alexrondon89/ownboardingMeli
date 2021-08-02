package coingecko_service

import (
	"log"
	"ownboardingMeli/internal/client"
	"ownboardingMeli/internal/platform"
	"ownboardingMeli/internal/service/dto"
	"sync"
)

type CoinGeckoService struct {
	CoinGeckoClient client.ClientInterface
}

func NewCoinGeckoService(client client.ClientInterface) *CoinGeckoService{
	return &CoinGeckoService{CoinGeckoClient: client }
}

func (s *CoinGeckoService) GetPrice(coin string, currency string) (*dto.CryptoResponse, error){
	clientResponse , err := s.CoinGeckoClient.GetPriceFromClient(coin, currency)
	if err != nil {
		return platform.BuildPartialResponse(coin), err
	}

	response := platform.BuildCryptoResponse(clientResponse, currency)
	return response, nil
}

func (s *CoinGeckoService) GetListPrice(coins []string, currency string) ([]dto.CryptoResponse, error){
	n := len(coins)
	var ListResponse []dto.CryptoResponse
	channel:= make(chan dto.CryptoResponse, n) //concurrency
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

func (s *CoinGeckoService) ReadCoinPrices(coin string, currency string, wg *sync.WaitGroup, c chan <- dto.CryptoResponse){
	defer wg.Done()
	defer s.Recover(coin, c)

	clientResponse , err := s.CoinGeckoClient.GetPriceFromClient(coin, currency)

	// Only use to test Recover method after a panic
	if coin == "bitcoin"{
		panic("bitcoin error")
	}

	if err != nil {
		c<- *platform.BuildPartialResponse(coin)
	}

	c<- *platform.BuildCryptoResponse(clientResponse, currency)
}

func (s *CoinGeckoService) Recover(coin string, c chan <- dto.CryptoResponse){
	if r := recover(); r!= nil{
		log.Println("panic occurred: ",r)
		c <- *platform.BuildPartialResponse(coin)
	}
}
