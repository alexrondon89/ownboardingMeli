package service

type CryptoService interface {
	GetPrice(id string, currency string) (*CryptoResponse, error)
	GetListPrice(coins []string, currency string) ([]CryptoResponse, error)
}

type CryptoResponse struct {
	Id        	string      `json:"id"`
	Content		*Content	`json:"content,omitempty"`
	Partial 	bool		`json:"partial"`
}

type Content struct {
	Price		float64	`json:"price"`
	Currency	string		`json:"currency"`
}

type ListCryptoResponse struct {
	Items []CryptoResponse		`json:"items"`
}