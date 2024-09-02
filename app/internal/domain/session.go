package domain

import "app/internal/lib"

type (
	SessionID string
)

func NewSessionID() SessionID {
	return SessionID(lib.RandomString(10))
}

func (s SessionID) String() string {
	return string(s)
}

func (s SessionID) ToCookieKey() string {
	return "session_id"
}
