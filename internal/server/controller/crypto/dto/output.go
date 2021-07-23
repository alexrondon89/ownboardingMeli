package dto

type PartialResponse struct {
	Id		string 	`json:"id"`
	Partial	bool	`json:"partial"`
}

func BuildPartialResponse(id string) *PartialResponse{
	return &PartialResponse{Id: id, Partial: true}
}