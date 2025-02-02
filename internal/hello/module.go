package hello

import (
	"eventsguard/internal/hello/adapters/api"

	"go.uber.org/fx"
)

var Module = fx.Module("hello",
	api.Module,
)
