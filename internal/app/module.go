// internal/app/module.go.
package app

import (
	"eventsguard/internal/hello"

	"go.uber.org/fx"
)

var Module = fx.Module("app",
	hello.Module,
)
