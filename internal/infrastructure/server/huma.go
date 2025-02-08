package server

import (
	"context"
	"fmt"
	"net/http"

	"eventsguard/internal/infrastructure/config"

	"eventsguard/internal/utils/fx/route"
	"eventsguard/internal/utils/huma/format"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
)

type humaServer struct {
	cfg    *config.AppConfig
	server Server
	routes []route.Route
}

func NewHumaServer(
	cfg *config.AppConfig,
	server Server,
	routes []route.Route,

) Server {

	Formats := huma.DefaultFormats
	Formats["application/json"] = format.DefaultJSONFormat
	apiConfig := huma.DefaultConfig(
		cfg.ApiName,
		cfg.ApiVersion,
	)
	apiConfig.Formats = Formats
	apiConfig.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"TokenAuth": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
		"Admin": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}

	url := fmt.Sprintf("%s/api", cfg.ServerUri)
	apiConfig.Servers = []*huma.Server{{URL: url}}
	api := humago.NewWithPrefix(server.GetMux(), "/api", apiConfig)

	// huma.Get(api, "/demo/", func(ctx context.Context, input *struct{}) (*struct{}, error) {
	// 	return nil, nil
	// })

	for _, route := range routes {
		route.Register(api)
	}
	humaServer := &humaServer{
		cfg:    cfg,
		server: server,
		routes: routes,
	}
	return humaServer
}

func (s *humaServer) Start() {
	s.server.Start()
}

func (s *humaServer) GetMux() *http.ServeMux {
	return s.server.GetMux()
}

func (s *humaServer) Shutdown(ctx context.Context) error {
	return nil
}
