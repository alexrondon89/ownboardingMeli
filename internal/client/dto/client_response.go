package dto

type ClientResponse struct {
	Id			string
	Price 		float64
	Currency 	string
}

func NewClientResponse() *ClientResponse {
	return &ClientResponse{}
}