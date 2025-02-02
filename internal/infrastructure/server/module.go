package server

import (
	"context"

	"go.uber.org/fx"
)

var Module = fx.Module("server",
	fx.Provide(
		fx.Annotate(NewMuxServer, fx.As(new(Server)), fx.ResultTags(`name:"server"`)),
		fx.Annotate(NewHumaServer, fx.As(new(Server)), fx.ResultTags(`name:"api"`), fx.ParamTags("", `name:"server"`, `group:"routes"`)),
	),
	fx.Invoke(
		fx.Annotate(
			func(lc fx.Lifecycle, api Server) {
				lc.Append(fx.Hook{
					OnStart: func(ctx context.Context) error {
						go api.Start()
						return nil
					},
					OnStop: func(ctx context.Context) error {
						return nil
					},
				})
			},
			fx.ParamTags(``, `name:"api"`),
		),
	),
)
