package http

import (
	"net/http"

	"eventsguard/internal/utils/fx/route"

	"github.com/danielgtaylor/huma/v2"
)

type helloSimpleRouter struct {
	handler *helloSimpleHandler
}

func NewHelloSimpleRouter(
	handler *helloSimpleHandler,
) route.Route {
	return &helloSimpleRouter{
		handler,
	}
}

func (u helloSimpleRouter) Register(
	api huma.API,
) {

	huma.Register(api, huma.Operation{
		OperationID:   "hello-hello",
		Method:        http.MethodGet,
		Path:          "/hello/hello",
		Description:   "Simple Hello",
		Tags:          []string{"Hello"},
		DefaultStatus: http.StatusOK,
	}, u.handler.Handler)

}
