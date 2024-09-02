package handler

import (
	"app/internal/domain"
	"app/internal/handler/oapi"
	"app/internal/handler/util"
	"app/internal/lib"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
)

type Handlers struct {
	WebAuthn *webauthn.WebAuthn
}

func NewHandlers(wc *webauthn.WebAuthn) *Handlers {
	return &Handlers{
		WebAuthn: wc,
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
	// emailを@マークで分割してDisplayNameに設定
	emailParts := strings.Split(strEmail, "@")
	displayName := emailParts[0]
	user := &domain.User{
		ID:          []byte(lib.RandomString(20)),
		Name:        strEmail,
		DisplayName: displayName,
	}

	options, session, err := h.WebAuthn.BeginRegistration(user)
	if err != nil {
		util.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to begin registration")
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode("RegisterChallengePasskey")
}
