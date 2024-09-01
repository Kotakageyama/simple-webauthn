package handler

import (
	"app/internal/handler/oapi"
	"encoding/json"
	"net/http"
)

type Handlers struct{}

func NewHandlers() *Handlers {
	return &Handlers{}
}

func (h *Handlers) LoginPasskey(w http.ResponseWriter, r *http.Request, params oapi.LoginPasskeyParams) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("LoginPasskey")
}

func (h *Handlers) LoginChallengePasskey(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("LoginChallengePasskey")
}

func (h *Handlers) RegisterPasskey(w http.ResponseWriter, r *http.Request, params oapi.RegisterPasskeyParams) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("RegisterPasskey")
}

func (h *Handlers) RegisterChallengePasskey(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("RegisterChallengePasskey")
}
