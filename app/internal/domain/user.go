package domain

import (
	"app/internal/lib"
	"strings"

	"github.com/go-webauthn/webauthn/webauthn"
)

type User struct {
	ID          []byte `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"display_name,omitempty"`

	Credentials []webauthn.Credential `json:"-"`
}

func NewUser(email string) User {
	// emailを@マークで分割してDisplayNameに設定
	emailParts := strings.Split(email, "@")
	displayName := emailParts[0]
	return User{
		ID:          []byte(lib.RandomString(20)),
		Name:        email,
		DisplayName: displayName,
	}
}

func (u *User) WebAuthnID() []byte {
	return u.ID
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	if u.DisplayName != "" {
		return u.DisplayName
	}
	return u.Name
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (u *User) WebAuthnIcon() string {
	return ""
}
