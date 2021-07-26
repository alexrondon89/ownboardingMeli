package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"ownboardingMeli/pkg/http/dto"
)

func GetRequest(url string) (*dto.HttpResponse, error){
	resp, err := http.Get(url)
	log.Println("imprime sl status code")
	log.Println(resp.StatusCode)

	if err != nil {
		return nil, err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := &dto.HttpResponse{Body: body, StatusCode: resp.StatusCode}
	return response,nil
}
