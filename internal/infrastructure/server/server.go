package server

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
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
	slog.Info("HTTP server starting on %s", s.server.Addr)

	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		slog.Error("Failed to start HTTP server: %v", err)
		return err
	}
	return nil
}

func (s *Server) Stop() error {
	slog.Info("Shutting down HTTP server on %s", s.server.Addr)
	return s.server.Shutdown(context.Background())
}

// WaitForSignal blocks until a signal is received
func (s *Server) WaitForSignal() {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan
	if err := s.Stop(); err != nil {
		slog.Error("Failed to stop server gracefully: %v", err)
	}
}
