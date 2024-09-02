package usecase

import (
	"app/internal/domain"
	"app/internal/repository"
	"fmt"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
)

type LoginUsecase interface {
	LoginChallenge() (*protocol.CredentialAssertion, domain.SessionID, error)
	LoginPasskey(sessionID domain.SessionID, request *protocol.ParsedCredentialAssertionData) (domain.User, error)
}

type loginUsecase struct {
	session repository.SessionRepository
	user    repository.UserRepository
	webAuth *webauthn.WebAuthn
}

func NewLoginUsecase(
	session repository.SessionRepository,
	user repository.UserRepository,
	webauthn *webauthn.WebAuthn,
) LoginUsecase {
	return &loginUsecase{
		session: session,
		user:    user,
		webAuth: webauthn,
	}
}

func (u *loginUsecase) LoginChallenge() (*protocol.CredentialAssertion, domain.SessionID, error) {
	credential, session, err := u.webAuth.BeginDiscoverableLogin()
	if err != nil {
		return nil, "", fmt.Errorf("failed to begin discoverable login: %w", err)
	}

	sessionID, err := u.session.Insert(session)
	if err != nil {
		return nil, "", fmt.Errorf("failed to insert login session: %w", err)
	}

	return credential, sessionID, nil
}

func (u *loginUsecase) LoginPasskey(
	sessionID domain.SessionID,
	request *protocol.ParsedCredentialAssertionData,
) (domain.User, error) {
	session, err := u.session.Get(sessionID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get session: %w", err)
	}

	var loggedInUser *domain.User = nil
	handler := func(rawID, userHandle []byte) (webauthn.User, error) {
		user, err := u.user.GetByUserId(userHandle)
		if err != nil {
			return nil, fmt.Errorf("failed to get user by user ID: %w", err)
		}
		loggedInUser = user
		return user, nil
	}

	_, err = u.webAuth.ValidateDiscoverableLogin(handler, *session, request)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to create credential: %w", err)
	}

	return *loggedInUser, nil
}
