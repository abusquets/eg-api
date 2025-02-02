package server

import (
	"context"
	"eventsguard/internal/infrastructure/config"
	"fmt"

	"log"
	"net/http"

	"github.com/rs/cors"
)

type muxServer struct {
	config  *config.AppConfig
	mux     *http.ServeMux
	handler http.Handler
}

func index(w http.ResponseWriter, r *http.Request) {
	if _, err := w.Write([]byte("oli Mux, ke ase?")); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func NewMuxServer(cfg *config.AppConfig) Server {
	mux := http.NewServeMux()

	mux.Handle("GET /", http.HandlerFunc(index))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})

	return &muxServer{
		config:  cfg,
		mux:     mux,
		handler: c.Handler(mux),
	}

}

func (s *muxServer) Start() {
	log.Println("Starting the server...")

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", s.config.ServerPort),
		Handler: s.handler,
	}

	if err := httpServer.ListenAndServe(); err != nil {
		log.Printf("Server failed to start: %v", err)
	}
}

func (s *muxServer) Shutdown(ctx context.Context) error {
	log.Println("Shutting down the server gracefully...")
	return nil
}

func (s *muxServer) GetMux() *http.ServeMux {
	return s.mux
}
