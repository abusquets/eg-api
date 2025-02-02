package api

import (
	"eventsguard/internal/hello/adapters/api/http"

	"go.uber.org/fx"
)

var Module = fx.Options(
	http.Module,
)
