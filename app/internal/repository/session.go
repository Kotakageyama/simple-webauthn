package repository

import (
	"app/internal/domain"

	"github.com/go-webauthn/webauthn/webauthn"
	"golang.org/x/xerrors"
)

type SessionRepository interface {
	Insert(session *webauthn.SessionData) (domain.SessionID, error)
	Get(id domain.SessionID) (*webauthn.SessionData, error)
}

type sessionRepository struct {
	Sessions map[domain.SessionID]*webauthn.SessionData
}

func NewSessionRepository() SessionRepository {
	return &sessionRepository{
		Sessions: map[domain.SessionID]*webauthn.SessionData{},
	}
}

func (r *sessionRepository) Insert(session *webauthn.SessionData) (domain.SessionID, error) {
	id := domain.NewSessionID()
	if _, exists := r.Sessions[id]; exists {
		return "", xerrors.New("session already exists")
	}
	r.Sessions[id] = session
	return id, nil
}

func (r *sessionRepository) Get(id domain.SessionID) (*webauthn.SessionData, error) {
	s, ok := r.Sessions[id]
	if !ok {
		return nil, xerrors.New("session not found")
	}
	return s, nil
}
