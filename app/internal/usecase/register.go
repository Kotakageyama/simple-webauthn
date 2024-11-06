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
	options, waSession, err := u.webAuth.BeginRegistration(&user)
	if err != nil {
		return nil, "", err
	}

	// Convert WebAuthn session to domain session
	session := repository.ConvertWebAuthnSession(waSession, user.ID)

	err = u.session.Insert(session)
	if err != nil {
		return nil, "", fmt.Errorf("failed to insert session: %w", err)
	}

	err = u.user.Create(&user)
	if err != nil {
		return nil, "", fmt.Errorf("failed to insert user: %w", err)
	}

	return options, session.ID, nil
}

func (u *registerUsecase) RegisterPasskey(
	sessionID domain.SessionID,
	request *protocol.ParsedCredentialCreationData,
) error {
	session, err := u.session.Get(sessionID)
	if err != nil {
		return fmt.Errorf("failed to get session: %w", err)
	}

	user, err := u.user.Get(session.UserID)
	if err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	// Convert domain session back to WebAuthn session
	waSession := &webauthn.SessionData{
		Challenge:           session.WebAuthnData.Challenge,
		UserID:             session.WebAuthnData.UserID,
		AllowedCredentialIDs: session.WebAuthnData.AllowCredentials,
		UserVerification:    session.WebAuthnData.UserVerification,
		Extensions:         session.WebAuthnData.Extensions,
	}

	credential, err := u.webAuth.CreateCredential(user, *waSession, request)
	if err != nil {
		return fmt.Errorf("failed to create credential: %w", err)
	}

	user.Credentials = append(user.Credentials, *credential)
	err = u.user.Update(user)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
