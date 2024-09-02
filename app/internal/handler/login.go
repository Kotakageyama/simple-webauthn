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
	login    usecase.LoginUsecase
}

func NewHandlers(register usecase.RegisterUsecase, login usecase.LoginUsecase) *Handlers {
	return &Handlers{
		register: register,
		login:    login,
	}
}

func (h *Handlers) LoginPasskey(w http.ResponseWriter, r *http.Request) {
	// クッキーからセッションIDを取得
	cookie, err := r.Cookie("session_id")
	if err != nil {
		util.WriteErrorResponse(w, http.StatusUnauthorized, "Session ID not found", err)
		return
	}
	sessionID := cookie.Value

	request, err := protocol.ParseCredentialRequestResponseBody(r.Body)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	user, err := h.login.LoginPasskey(domain.SessionID(sessionID), request)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to login passkey", err)
		return
	}
	util.WriteResponse(w, http.StatusOK, user)
}

func (h *Handlers) LoginChallengePasskey(w http.ResponseWriter, r *http.Request) {
	credential, sessionID, err := h.login.LoginChallenge()
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to begin login", err)
		return
	}
	util.SetCookie(w, sessionID.ToCookieKey(), sessionID.String())
	util.WriteResponse(w, http.StatusOK, credential)
}

func (h *Handlers) RegisterPasskey(w http.ResponseWriter, r *http.Request) {
	// クッキーからセッションIDを取得
	cookie, err := r.Cookie("session_id")
	if err != nil {
		util.WriteErrorResponse(w, http.StatusUnauthorized, "Session ID not found", err)
		return
	}
	sessionID := cookie.Value

	request, err := protocol.ParseCredentialCreationResponseBody(r.Body)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	err = h.register.RegisterPasskey(domain.SessionID(sessionID), request)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to register passkey", err)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) RegisterChallengePasskey(w http.ResponseWriter, r *http.Request) {
	var params oapi.RegisterChallengePasskeyRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusBadRequest, "Invalid request payload", err)
		return
	}

	strEmail := string(params.Email)

	options, sessionID, err := h.register.RegisterChallenge(strEmail)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to begin registration", err)
		return
	}
	util.SetCookie(w, sessionID.ToCookieKey(), sessionID.String())
	util.WriteResponse(w, http.StatusOK, options)
}
