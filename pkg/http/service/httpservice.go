package service

import (
	"io/ioutil"
	httpcli "net/http"
	"ownboardingMeli/pkg/http"
)

func GetRequest(url string) (*http.HttpResponse, error){
	resp, err := httpcli.Get(url)

	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := http.NewHttpResponse(body, resp.StatusCode)
	return response,nil
}

func CheckStatusCode200 (code int) bool{
	return code == 200
}
