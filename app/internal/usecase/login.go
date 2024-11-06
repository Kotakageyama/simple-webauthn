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
	LoginWorldID(nullifier string, proof string) (domain.User, domain.SessionID, error)
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
	credential, waSession, err := u.webAuth.BeginDiscoverableLogin()
	if err != nil {
		return nil, "", fmt.Errorf("failed to begin discoverable login: %w", err)
	}

	// Create a new session with WebAuthn data
	session := repository.ConvertWebAuthnSession(waSession, nil)

	if err := u.session.Insert(session); err != nil {
		return nil, "", fmt.Errorf("failed to insert login session: %w", err)
	}

	return credential, session.ID, nil
}

func (u *loginUsecase) LoginPasskey(
	sessionID domain.SessionID,
	request *protocol.ParsedCredentialAssertionData,
) (domain.User, error) {
	session, err := u.session.Get(sessionID)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to get session: %w", err)
	}

	if session.AuthMethod != "passkey" {
		return domain.User{}, fmt.Errorf("invalid authentication method for session")
	}

	// Convert back to WebAuthn session data
	waSession := &webauthn.SessionData{
		Challenge:           session.WebAuthnData.Challenge,
		UserID:             session.WebAuthnData.UserID,
		AllowedCredentialIDs: session.WebAuthnData.AllowCredentials,
		UserVerification:    session.WebAuthnData.UserVerification,
		Extensions:         session.WebAuthnData.Extensions,
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

	_, err = u.webAuth.ValidateDiscoverableLogin(handler, *waSession, request)
	if err != nil {
		return domain.User{}, fmt.Errorf("failed to validate login: %w", err)
	}

	// Update session with user information
	session.UserID = loggedInUser.ID
	if err := u.session.Update(session); err != nil {
		return domain.User{}, fmt.Errorf("failed to update session: %w", err)
	}

	return *loggedInUser, nil
}

func (u *loginUsecase) LoginWorldID(nullifier string, proof string) (domain.User, domain.SessionID, error) {
	// Find or create user based on World ID nullifier
	user, err := u.user.GetByWorldIDNullifier(nullifier)
	if err != nil {
		// Create new user if not found
		user = &domain.User{
			ID:              []byte(nullifier), // Use nullifier as ID for World ID users
			AuthMethod:      "worldid",
			WorldIDVerified: true,
		}
		if err := u.user.Create(user); err != nil {
			return domain.User{}, "", fmt.Errorf("failed to create user: %w", err)
		}
	}

	// Create new session for World ID authentication
	session := domain.NewSession(user.ID, "worldid")
	if err := u.session.Insert(session); err != nil {
		return domain.User{}, "", fmt.Errorf("failed to create session: %w", err)
	}

	return *user, session.ID, nil
}
