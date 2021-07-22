package models

type Input struct {
	Data string `json:"data" binding:"required"`
}

type Response struct {
	Message	string `json:"message"`
}

