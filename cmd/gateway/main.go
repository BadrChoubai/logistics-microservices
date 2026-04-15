package main

import (
	"context"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/BadrChoubai/logistics-microservices/internal/config"
	"github.com/BadrChoubai/logistics-microservices/internal/observability/logger"
	"github.com/BadrChoubai/logistics-microservices/internal/server"
)

type application struct {
	config *config.Gateway
	Log    *logger.Logger
}

// @title						Logistics Services API Gateway
// @version					1.0.0
// @description				API Gateway for the Logistics Services API platform
// @host						localhost:8080
// @BasePath					/api/v1
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
		cfgPath = "manifests/gateway/config.json"
	}

	cfg, err := config.Load[*config.Gateway](cfgPath)
	if err != nil {
		return err
	}

	app := &application{
		config: cfg,
		Log:    logger.NewLogger(stdout, cfg.LogLevel),
	}

	srv, err := server.NewServer(app.config.Port, app.Log)

	if err != nil {
		return err
	}

	// Create a context that cancels on SIGINT or SIGTERM
	signalCtx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Run the server in a goroutine
	serverErr := make(chan error, 1)
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			serverErr <- err
			return
		}
		serverErr <- nil
	}()

	// Wait for shutdown signal
	<-signalCtx.Done()
	app.Log.Info("Shutdown signal received")

	// Perform graceful shutdown with timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		return err
	}

	// Check if server encountered other errors
	if err := <-serverErr; err != nil {
		return err
	}

	return nil
}
