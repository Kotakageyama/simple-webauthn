package domain

import (
	"app/internal/lib"
	"github.com/go-webauthn/webauthn/protocol"
)

type (
	SessionID string
)

type Session struct {
	ID         SessionID
	UserID     []byte
	AuthMethod string // "passkey" or "worldid"
	WebAuthnData *WebAuthnSessionData
}

type WebAuthnSessionData struct {
	Challenge        string
	UserID          []byte
	AllowCredentials [][]byte
	UserVerification protocol.UserVerificationRequirement
	Extensions      map[string]interface{}
}

func NewSessionID() SessionID {
	return SessionID(lib.RandomString(10))
}

func NewSession(userID []byte, authMethod string) *Session {
	return &Session{
		ID:         NewSessionID(),
		UserID:     userID,
		AuthMethod: authMethod,
	}
}

func (s SessionID) String() string {
	return string(s)
}

func (s SessionID) ToCookieKey() string {
	return "session_id"
}
