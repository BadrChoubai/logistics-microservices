package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/BadrChoubai/logistics-microservices/api/swagger/shipment"
	"github.com/BadrChoubai/logistics-microservices/internal/shipment/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title						Shipment Service
// @version					1.0.0
// @description				Shipment Service for the Logistics Services API platform
// @host						localhost:8081
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
		httpSwagger.URL("http://localhost:8081/swagger/doc.json"),
	))
	mux.HandleFunc("/health", handler.GetShipmentHealth)
	fmt.Println("server running on localhost:8081")

	err := http.ListenAndServe(":8081", mux)
	if err != nil {
		return err
	}

	return nil
}
