package server

import (
	"context"
	"errors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	server *http.Server
}

func NewServer(handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr: os.Getenv("HTTP_SERVER_ADDRESS"),
			// Use h2c for HTTP/2 without TLS
			Handler: h2c.NewHandler(handler, &http2.Server{}),
		},
	}
}

func (s *Server) Start() error {
	log.Printf("HTTP server starting on %s", s.server.Addr)

	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("Failed to start HTTP server: %v", err)
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	log.Printf("Shutting down HTTP server on %s", s.server.Addr)
	return s.server.Shutdown(context.Background())
}

// WaitForSignal blocks until a signal is received
func (s *Server) WaitForSignal() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	if err := s.Stop(); err != nil {
		log.Fatalf("Failed to stop server gracefully: %v", err)
	}
}
