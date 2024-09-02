//go:build wireinject
// +build wireinject

package di

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/usecase"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/wire"
	"gorm.io/gorm"
)

var repositorySet = wire.NewSet(
	repository.NewSessionRepository,
	repository.NewUserRepository,
)

var usecaseSet = wire.NewSet(
	usecase.NewRegisterUsecase,
)

var handlerSet = wire.NewSet(
	handler.NewHandlers,
)

func Wire(
	db *gorm.DB,
	WebAuthn *webauthn.WebAuthn,
) *handler.Handlers {
	wire.Build(
		repositorySet,
		usecaseSet,
		handlerSet,
	)
	return &handler.Handlers{}
}
