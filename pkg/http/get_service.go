package http

import (
	"io/ioutil"
	"net/http"
	"ownboardingMeli/pkg/http/dto"
)

func GetRequest(url string) (*dto.HttpResponse, error){
	resp, err := http.Get(url)

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
