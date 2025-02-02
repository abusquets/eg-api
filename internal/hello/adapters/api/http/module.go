package http

import (
	"eventsguard/internal/utils/fx/route"

	"go.uber.org/fx"
)

var Module = fx.Options(
	fx.Provide(
		NewHelloSimpleHandler,
		route.AsRoute(NewHelloSimpleRouter),
	),
)
