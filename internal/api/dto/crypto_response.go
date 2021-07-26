package dto

type CryptoResponse struct {
	Id        	string      `json:"id"`
	Content		Content		`json:"content"`
	Partial 	bool		`json:"partial"`
}

type Content struct {
	Price		float64		`json:"price"`
	Currency	string		`json:"currency"`
}

type ListCryptoResponse struct {
	Items []CryptoResponse		`json:"items"`
}

type PartialResponse struct {
	Id		string 	`json:"id"`
	Partial	bool	`json:"partial"`
}


func BuildPartialResponse(id string) *PartialResponse{
	return &PartialResponse{Id: id, Partial: true}
}

