package util

import (
	"encoding/json"
	"log"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func WriteErrorResponse(w http.ResponseWriter, code int, message string, err error) {
	log.Printf("Message: %v, Error: %+v", message, err)
	w.WriteHeader(code)
	response := ErrorResponse{
		Code:    code,
		Message: message,
	}
	_ = json.NewEncoder(w).Encode(response)
}

func WriteResponse(w http.ResponseWriter, code int, data any) {
	w.WriteHeader(code)
	_ = json.NewEncoder(w).Encode(data)
}

func SetCookie(w http.ResponseWriter, name, value string) {
	http.SetCookie(w, &http.Cookie{
		Name:  name,
		Value: value,
	})
}
