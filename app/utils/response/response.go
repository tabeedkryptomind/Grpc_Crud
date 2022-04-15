package response

import (
	"net/http"
)

type ResRespo struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func ReuqstAccepted(message string, ) *ResRespo {
	return &ResRespo{
		Message: message,
		Status:  http.StatusOK,
	}
}

