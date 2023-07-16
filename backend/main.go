package main

import (
	"github.com/sidnarsipur/protect-ai/bootstrap"
	"github.com/sidnarsipur/protect-ai/lib"
	"go.uber.org/fx"
)

func main() {
	logger := lib.GetLogger().GetFxLogger()
	fx.New(bootstrap.Module, fx.Logger(logger)).Run()
}
