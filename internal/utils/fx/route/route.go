package route

import (
	"github.com/danielgtaylor/huma/v2"
	"go.uber.org/fx"
)

type Route interface {
	Register(api huma.API)
}

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
