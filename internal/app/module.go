// internal/app/module.go.
package app

import (
	"eventsguard/internal/infrastructure/config"

	"go.uber.org/fx"
)

var Module = fx.Module("app",
	config.Module,
)
