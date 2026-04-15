package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BadrChoubai/logistics-microservices/cmd/inventory/server"
	"github.com/BadrChoubai/logistics-microservices/internal/config"
	"github.com/BadrChoubai/logistics-microservices/internal/observability/logger"
)

// @title						Inventory Service
// @version					1.0.0
// @description				Inventory Service for the Logistics Services API platform
// @host						localhost:8082
// @BasePath					/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	ctx := context.Background()

	if err := run(ctx, os.Stdout, os.Getenv); err != nil {
		log.Fatalln(err)
	}
}

func run(ctx context.Context, stdout io.Writer, getenv func(string) string) error {
	// CONFIG_PATH can be overridden per environment (Docker, K8s, local).
	cfgPath := getenv("CONFIG_PATH")
	if cfgPath == "" {
		cfgPath = "manifests/inventory/config.json"
	}

	cfg, err := config.Load[*config.Service](cfgPath)
	if err != nil {
		return err
	}

	logger := logger.NewLogger(stdout, cfg.LogLevel)
	srv, err := server.NewServer(cfg.Port, logger)

	if err != nil {
		return err
	}

	shutdownError := make(chan error)

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		s := <-quit

		logger.Info(fmt.Sprintf("caught signal: %s", s))

		shutdownCtx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		err := srv.Shutdown(shutdownCtx)
		if err != nil {
			shutdownError <- err
		}

		shutdownError <- nil
	}()

	err = srv.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	err = <-shutdownError
	if err != nil {
		return err
	}

	return nil
}
