//go:build wireinject
// +build wireinject

package di

import (
	"app/internal/handler"

	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/google/wire"
	"gorm.io/gorm"
)

// プロバイダーセットの定義
var handlerSet = wire.NewSet(
	handler.NewHandlers,
)

func Wire(
	db *gorm.DB,
	WebAuthn *webauthn.WebAuthn,
) *handler.Handlers {
	wire.Build(
		handlerSet,
	)
	return &handler.Handlers{}
}
