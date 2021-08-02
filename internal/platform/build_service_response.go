//	methods that can be use across internal folder for different services
package platform

import (
	client "ownboardingMeli/internal/client/dto"
	service "ownboardingMeli/internal/service/dto"
	"strings"
)

func BuildCryptoResponse(clientResponse *client.ClientResponse, currency string) *service.CryptoResponse{
	return &service.CryptoResponse{
		Id: clientResponse.Id,
		Content: &service.Content{
			Price: clientResponse.Price,
			Currency: strings.ToUpper(currency),
		},
		Partial: false,
	}
}

func BuildPartialResponse(id string) *service.CryptoResponse{
	return &service.CryptoResponse{
		Id: id,
		Partial: true,
	}
}
