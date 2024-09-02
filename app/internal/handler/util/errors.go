package util

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteErrorResponse(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	response := ErrorResponse{
		Code:    code,
		Message: message,
	}
	_ = json.NewEncoder(w).Encode(response)
}
