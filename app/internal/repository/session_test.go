package repository

import (
	"app/internal/domain"
	"testing"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/stretchr/testify/assert"
)

func TestSessionRepository(t *testing.T) {
	repo := NewSessionRepository()

	t.Run("test passkey session", func(t *testing.T) {
		// Create a WebAuthn session
		waSession := &webauthn.SessionData{
			Challenge:           "test-challenge",
			UserID:             []byte("test-user"),
			AllowedCredentialIDs: [][]byte{[]byte("cred1"), []byte("cred2")},
			UserVerification:    protocol.VerificationRequired,
			Extensions:         map[string]interface{}{"ext1": "value1"},
		}

		// Convert and store session
		session := ConvertWebAuthnSession(waSession, []byte("test-user"))
		err := repo.Insert(session)
		assert.NoError(t, err)

		// Retrieve and verify
		retrieved, err := repo.Get(session.ID)
		assert.NoError(t, err)
		assert.Equal(t, "passkey", retrieved.AuthMethod)
		assert.Equal(t, waSession.Challenge, retrieved.WebAuthnData.Challenge)
		assert.Equal(t, waSession.UserID, retrieved.WebAuthnData.UserID)
		assert.Equal(t, waSession.AllowedCredentialIDs, retrieved.WebAuthnData.AllowCredentials)
		assert.Equal(t, waSession.UserVerification, retrieved.WebAuthnData.UserVerification)
	})

	t.Run("test world id session", func(t *testing.T) {
		// Create a World ID session
		session := domain.NewSession([]byte("world-id-user"), "worldid")
		err := repo.Insert(session)
		assert.NoError(t, err)

		// Retrieve and verify
		retrieved, err := repo.Get(session.ID)
		assert.NoError(t, err)
		assert.Equal(t, "worldid", retrieved.AuthMethod)
		assert.Equal(t, []byte("world-id-user"), retrieved.UserID)
		assert.Nil(t, retrieved.WebAuthnData)
	})

	t.Run("test session update", func(t *testing.T) {
		session := domain.NewSession([]byte("test-user"), "worldid")
		err := repo.Insert(session)
		assert.NoError(t, err)

		// Update session
		session.UserID = []byte("updated-user")
		err = repo.Update(session)
		assert.NoError(t, err)

		// Verify update
		retrieved, err := repo.Get(session.ID)
		assert.NoError(t, err)
		assert.Equal(t, []byte("updated-user"), retrieved.UserID)
	})

	t.Run("test session delete", func(t *testing.T) {
		session := domain.NewSession([]byte("test-user"), "worldid")
		err := repo.Insert(session)
		assert.NoError(t, err)

		// Delete session
		err = repo.Delete(session.ID)
		assert.NoError(t, err)

		// Verify deletion
		_, err = repo.Get(session.ID)
		assert.Error(t, err)
	})
}
