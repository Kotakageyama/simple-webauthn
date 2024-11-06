package usecase

import (
	"testing"

	"app/internal/repository"
	"github.com/stretchr/testify/assert"
)

func TestLoginWorldID(t *testing.T) {
	// Setup repositories
	sessionRepo := repository.NewSessionRepository()
	userRepo := repository.NewUserRepository()

	// Create login usecase
	loginUC := NewLoginUsecase(sessionRepo, userRepo, nil)

	tests := []struct {
		name      string
		nullifier string
		proof     string
		wantErr   bool
	}{
		{
			name:      "successful login with new user",
			nullifier: "test-nullifier",
			proof:     "test-proof",
			wantErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, sessionID, err := loginUC.LoginWorldID(tt.nullifier, tt.proof)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, sessionID)
			assert.Equal(t, "worldid", user.AuthMethod)
			assert.True(t, user.WorldIDVerified)
			assert.Equal(t, []byte(tt.nullifier), user.ID)

			// Verify session was created
			session, err := sessionRepo.Get(sessionID)
			assert.NoError(t, err)
			assert.Equal(t, "worldid", session.AuthMethod)
			assert.Equal(t, user.ID, session.UserID)
		})
	}
}
