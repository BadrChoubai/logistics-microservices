package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/BadrChoubai/logistics-microservices/api/swagger/gateway"
	"github.com/BadrChoubai/logistics-microservices/internal/gateway/handler"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title						Logistics Services API Gateway
// @version					1.0.0
// @description				API Gateway for the Logistics Services API platform
// @host						localhost:8080
// @BasePath					/api/v1
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
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))
	mux.HandleFunc("/health", handler.Health)

	fmt.Println("server running on localhost:8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		return err
	}

	return nil
}
