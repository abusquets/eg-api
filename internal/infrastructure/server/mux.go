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
	config *config.AppConfig
	mux    *http.ServeMux
}

// func index(w http.ResponseWriter, r *http.Request) {
// 	if _, err := w.Write([]byte("oli Mux, ke ase?")); err != nil {
// 		log.Printf("Failed to write response: %v", err)
// 	}
// }

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Healthy")
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func NewMuxServer(cfg *config.AppConfig) Server {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", index)
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/", fs)

	mux.HandleFunc("/health", health)

	return &muxServer{
		config: cfg,
		mux:    mux,
	}

}

func (s *muxServer) Start() {
	log.Println("Starting the server...")

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	handler := c.Handler(s.mux)

	port := fmt.Sprintf(":%s", s.config.ServerPort)
	if err := http.ListenAndServe(port, handler); err != nil {
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
