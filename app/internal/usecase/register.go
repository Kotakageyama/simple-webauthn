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
	RegisterPasskey(sessionID domain.SessionID, request *protocol.ParsedCredentialCreationData) error
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
	err = u.user.Create(sessionID, &user)
	if err != nil {
		return nil, "", fmt.Errorf("failed to insert user: %w", err)
	}

	return options, sessionID, nil
}

func (u *registerUsecase) RegisterPasskey(
	sessionID domain.SessionID,
	request *protocol.ParsedCredentialCreationData,
) error {
	user, err := u.user.Get(sessionID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}
	session, err := u.session.Get(sessionID)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}

	credential, err := u.webAuth.CreateCredential(user, *session, request)
	if err != nil {
		return fmt.Errorf("failed to create credential: %w", err)
	}

	user.Credentials = append(user.Credentials, *credential)
	err = u.user.Update(sessionID, user)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
