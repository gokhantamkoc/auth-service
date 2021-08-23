package util

import (
	"net/http"
)

type Response struct {
	Success bool 	`json:"success"`
	Code	int		`json:"code"`
	Message string	`json:"message"`
}

type SuccessResponse struct {
	Response
	Data 	interface{}		`json:"data"`
}

type ErrorResponse struct {
	Response
}

func CreateSuccessResponse(data interface{}) SuccessResponse {
	return SuccessResponse{
		Response: Response{
			Success: true,
			Code:    http.StatusOK,
			Message: "",
		},
		Data:     data,
	}
}

func CreateErrorResponse(statusCode int, message string) ErrorResponse {
	return ErrorResponse{
		Response: Response{
			Success: false,
			Code:    statusCode,
			Message: message,
		},
	}
}