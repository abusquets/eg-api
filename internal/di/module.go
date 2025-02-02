// internal/app/module.go.
package di

import (
	"eventsguard/internal/infrastructure/config"

	"go.uber.org/fx"
)

var BaseModule = fx.Module("base",
	config.Module,
)
