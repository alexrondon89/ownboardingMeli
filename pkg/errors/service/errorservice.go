package service

import "ownboardingMeli/pkg/errors"

const (
	BadRequest = iota
	InternalServer
	UnknownError
)

func BuildBadRequestError(message string) *errors.ErrorResponse {
	return &errors.ErrorResponse{Message: message, Code: BadRequest}
}

func BuildInternalServerError(message string) *errors.ErrorResponse {
	return &errors.ErrorResponse{Message: message, Code: InternalServer}
}

func BuildUnknownError(message string) *errors.ErrorResponse {
	return &errors.ErrorResponse{Message: message, Code: UnknownError}
}