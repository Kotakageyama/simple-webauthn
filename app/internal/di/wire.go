//go:build wireinject
// +build wireinject

package di

import (
	"app/internal/handler"

	"github.com/google/wire"
	"gorm.io/gorm"
)

// プロバイダーセットの定義
var handlerSet = wire.NewSet(
	handler.NewHandlers,
)

func Wire(
	db *gorm.DB,
) *handler.Handlers {
	wire.Build(
		handlerSet,
	)
	return &handler.Handlers{}
}
