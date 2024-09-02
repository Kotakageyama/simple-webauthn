package handler

import (
	"app/internal/domain"
	"app/internal/handler/oapi"
	"app/internal/handler/util"
	"app/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/go-webauthn/webauthn/protocol"
)

type Handlers struct {
	register usecase.RegisterUsecase
}

func NewHandlers(register usecase.RegisterUsecase) *Handlers {
	return &Handlers{
		register: register,
	}
}

func (h *Handlers) LoginPasskey(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("LoginPasskey")
}

func (h *Handlers) LoginChallengePasskey(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("LoginChallengePasskey")
}

func (h *Handlers) RegisterPasskey(w http.ResponseWriter, r *http.Request) {
	// クッキーからセッションIDを取得
	cookie, err := r.Cookie("session_id")
	if err != nil {
		util.WriteErrorResponse(w, http.StatusUnauthorized, "Session ID not found")
		return
	}
	sessionID := cookie.Value

	request, err := protocol.ParseCredentialCreationResponseBody(r.Body)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	err = h.register.RegisterPasskey(domain.SessionID(sessionID), request)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to register passkey")
		return
	}
	w.WriteHeader(http.StatusCreated)
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
