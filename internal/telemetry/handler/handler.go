package handler

import (
	"encoding/json"
	"net/http"
)

type HealthcheckResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type TelemetryReading struct{}

// Telemetry Service Health godoc
//
//	@Summary		Health check
//	@Description	Returns service health check response
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	HealthcheckResponse
//	@Router			/health [get]
//	@Router			/health [head]
func GetShipmentHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodHead {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// nolint
	json.NewEncoder(w).Encode(HealthcheckResponse{Status: "Healthy", Version: "1.0.0"})
}

// Get Sensor Reading godoc
//
//	@Summary		Get Sensor Reading
//	@Description	Returns a sensor reading
//	@Tags			telemetry
//	@Produce		json
//	@Success		200	{object}	TelemetryReading
//	@Router			/readings [get]
func GetTelemetry(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Ingest Sensor Reading godoc
//
//	@Summary		Ingest Sensor Reading
//	@Description	Returns a sensor reading
//	@Tags			telemetry
//	@Produce		json
//	@Success		201	{object}	TelemetryReading
//	@Router			/readings [post]
func UploadTelemetry(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Get Sensor Reading godoc
//
//	@Summary		Get sensor reading for a Shipment
//	@Description	Returns telemetry data for a given shipment
//	@Tags			telemetry
//	@Produce		json
//	@Success		200	{list}	TelemetryReading
//	@Router			/readings/{shipment_id} [get]
func GetTelemetryByShipmentID(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
