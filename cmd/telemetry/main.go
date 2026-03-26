package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/BadrChoubai/logistics-microservices/api/swagger/telemetry"
	"github.com/BadrChoubai/logistics-microservices/internal/telemetry/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title						Telemetry Service
// @version					1.0.0
// @description				Telemetry Service for the Logistics Services API platform
// @host						localhost:8083
// @BasePath					/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

func run() error {
	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8083/swagger/doc.json"),
	))
	mux.HandleFunc("/health", handler.GetShipmentHealth)
	fmt.Println("server running on localhost:8083")

	err := http.ListenAndServe(":8083", mux)
	if err != nil {
		return err
	}

	return nil
}
