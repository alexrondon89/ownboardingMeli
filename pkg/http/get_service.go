package http

import (
	"io/ioutil"
	"net/http"
)

func GetRequest(url string) ([]byte, error){
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body,nil
}
