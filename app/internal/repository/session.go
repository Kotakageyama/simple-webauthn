package repository

import (
	"app/internal/domain"
	"github.com/go-webauthn/webauthn/webauthn"
	"golang.org/x/xerrors"
)

type SessionRepository interface {
	Insert(session *domain.Session) error
	Get(id domain.SessionID) (*domain.Session, error)
	Update(session *domain.Session) error
	Delete(id domain.SessionID) error
}

type sessionRepository struct {
	Sessions map[domain.SessionID]*domain.Session
}

func NewSessionRepository() SessionRepository {
	return &sessionRepository{
		Sessions: map[domain.SessionID]*domain.Session{},
	}
}

func (r *sessionRepository) Insert(session *domain.Session) error {
	if _, exists := r.Sessions[session.ID]; exists {
		return xerrors.New("session already exists")
	}
	r.Sessions[session.ID] = session
	return nil
}

func (r *sessionRepository) Get(id domain.SessionID) (*domain.Session, error) {
	s, ok := r.Sessions[id]
	if !ok {
		return nil, xerrors.New("session not found")
	}
	return s, nil
}

func (r *sessionRepository) Update(session *domain.Session) error {
	if _, exists := r.Sessions[session.ID]; !exists {
		return xerrors.New("session not found")
	}
	r.Sessions[session.ID] = session
	return nil
}

func (r *sessionRepository) Delete(id domain.SessionID) error {
	if _, exists := r.Sessions[id]; !exists {
		return xerrors.New("session not found")
	}
	delete(r.Sessions, id)
	return nil
}

// Helper function to convert WebAuthn session data
func ConvertWebAuthnSession(waSession *webauthn.SessionData, userID []byte) *domain.Session {
	session := domain.NewSession(userID, "passkey")
	session.WebAuthnData = &domain.WebAuthnSessionData{
		Challenge:        waSession.Challenge,
		UserID:          waSession.UserID,
		AllowCredentials: waSession.AllowedCredentialIDs,
		UserVerification: waSession.UserVerification,
		Extensions:      waSession.Extensions,
	}
	return session
}
