package usecase

import (
	"app/internal/domain"
	"app/internal/repository"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type RegisterUsecase interface {
	RegisterChallenge(name string) (*protocol.CredentialCreation, domain.SessionID, error)
}

type registerUsecase struct {
	session repository.SessionRepository
	user    repository.UserRepository
	webAuth *webauthn.WebAuthn
}

func NewRegisterUsecase(
	session repository.SessionRepository,
	user repository.UserRepository,
	webauthn *webauthn.WebAuthn,
) RegisterUsecase {
	return &registerUsecase{
		session: session,
		user:    user,
		webAuth: webauthn,
	}
}

func (u *registerUsecase) RegisterChallenge(
	email string,
) (*protocol.CredentialCreation, domain.SessionID, error) {
	user := domain.NewUser(email)
	options, session, err := u.webAuth.BeginRegistration(&user)
	if err != nil {
		return nil, "", err
	}

	sessionID, err := u.session.Insert(session)
	if err != nil {
		return nil, "", fmt.Errorf("failed to insert session: %w", err)
	}
	err = u.user.Insert(sessionID, &user)
	if err != nil {
		return nil, "", fmt.Errorf("failed to insert user: %w", err)
	}

	return options, sessionID, nil
}
