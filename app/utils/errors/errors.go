package errors

import (
	"net/http"
)

type ResErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func InternalServerError(message string) *ResErr {
	return &ResErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "internal server error",
	}
}
func BadRequestError(message string) *ResErr {
	return &ResErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad request",
	}
}
