package server

import (
	"context"
	"net/http"
)

type Server interface {
	Start()
	GetMux() *http.ServeMux
	Shutdown(ctx context.Context) error
}
