package handler

import (
	"app/internal/handler/oapi"
	"app/internal/handler/util"
	"app/internal/usecase"
	"encoding/json"
	"net/http"
)

type Handlers struct {
	register usecase.RegisterUsecase
}

func NewHandlers(register usecase.RegisterUsecase) *Handlers {
	return &Handlers{
		register: register,
	}
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
	var params oapi.RegisterChallengePasskeyRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	strEmail := string(params.Email)

	options, sessionID, err := h.register.RegisterChallenge(strEmail)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to begin registration")
		return
	}
	util.SetCookie(w, sessionID.ToCookieKey(), sessionID.String())
	util.WriteResponse(w, http.StatusOK, options)
}
