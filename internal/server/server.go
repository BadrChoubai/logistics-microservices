// Package server
package server

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/BadrChoubai/logistics-microservices/internal/middleware"
	"github.com/BadrChoubai/logistics-microservices/internal/observability/logger"
)

var _ HTTPServer = (*Server)(nil)

type HTTPServer interface {
	ListenAndServe() error
	Shutdown(context.Context) error
}

type Server struct {
	HttpServer *http.Server
	Logger     *logger.Logger
}

func NewServer(port int, logger *logger.Logger) (*Server, error) {
	handler := setupHandler()

	srv := &Server{
		HttpServer: &http.Server{
			Addr:         ":" + strconv.Itoa(port),
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 5 * time.Second,
			IdleTimeout:  120 * time.Second,
			Handler:      handler,
		},
		Logger: logger,
	}

	return srv, nil
}

func setupHandler() http.Handler {
	mux := http.NewServeMux()

	var handler http.Handler = mux
	handler = middleware.Heartbeat(handler, "/health")

	return handler
}

func (s *Server) ListenAndServe() error {
	s.Logger.Info("Starting HTTP Server")
	s.Logger.Info("http://localhost" + s.HttpServer.Addr)

	return s.HttpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	s.Logger.Info("Shutting down HTTP Server")

	return s.HttpServer.Shutdown(ctx)
}
