package bootstrap

import (
	"context"

	"github.com/sidnarsipur/protect-ai/generate"
	"github.com/sidnarsipur/protect-ai/lib"
	"go.uber.org/fx"
)

// Module for initializing application
var Module = fx.Options(
	lib.Module,
	generate.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	logger lib.Logger,
) {

	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("---------------------")
			logger.Info("---------------------")
			logger.Info("Starting Backend")
			logger.Info("---------------------")
			logger.Info("---------------------")

			go func() {
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping Backend")
			return nil
		},
	})
}
