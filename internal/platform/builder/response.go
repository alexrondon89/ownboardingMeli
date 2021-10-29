//	methods that can be use across internal folder for different services
package builder

import (
	"ownboardingMeli/internal/client"
	"ownboardingMeli/internal/service"
	"strings"
)

func SuccessfulResponse(clientResponse *client.Response, currency string) *service.CryptoResponse{
	return &service.CryptoResponse{
		Id: clientResponse.Id,
		Content: &service.Content{
			Price: clientResponse.Price,
			Currency: strings.ToUpper(currency),
		},
		Partial: false,
	}
}

func PartialResponse(id string) *service.CryptoResponse{
	return &service.CryptoResponse{
		Id: id,
		Partial: true,
	}
}
