package Errors

import "ownboardingMeli/pkg/Errors/dto"

const (
	BadRequest = iota
	InternalServer
	UnknownError
)

func BuildBadRequestError(message string) *dto.ErrorResponse{
	return &dto.ErrorResponse{Message: message, Code: BadRequest}
}

func BuildInternalServerError(message string) *dto.ErrorResponse{
	return &dto.ErrorResponse{Message: message, Code: InternalServer}
}

func BuildUnknownError(message string) *dto.ErrorResponse{
	return &dto.ErrorResponse{Message: message, Code: UnknownError}
}