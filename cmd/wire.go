//go:build wireinject
// +build wireinject

package cmd

import (
	"context"

	"github.com/comeonjy/box/configs"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/google/wire"

	"github.com/comeonjy/box/internal/data"
	"github.com/comeonjy/box/internal/server"
	"github.com/comeonjy/box/internal/service"
)

func InitApp(ctx context.Context, logger *xlog.Logger) *App {
	panic(wire.Build(
		server.ProviderSet,
		service.ProviderSet,
		newApp,
		configs.ProviderSet,
		data.ProviderSet,
	))
}
