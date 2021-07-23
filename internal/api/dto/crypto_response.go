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

