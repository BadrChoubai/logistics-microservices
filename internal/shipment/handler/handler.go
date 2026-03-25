package handler

import (
	"encoding/json"
	"net/http"
)

type HealthcheckResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type ShipmentStatusResponse struct{}

type Shipment struct{}

// Shipment Service Health godoc
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

// Get Shipments godoc
//
//	@Summary		Get Shipments
//	@Description	Returns a list of shipments
//	@Tags			shipment
//	@Produce		json
//	@Success		200	{array}	Shipment
//	@Router			/shipments [get]
func GetShipments(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Create Shipment godoc
//
//	@Summary		Create Shipment
//	@Description	Create a new shipment
//	@Tags			shipment
//	@Produce		json
//	@Success		201	{object}	Shipment
//	@Router			/shipments [post]
func CreateShipment(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Get Shipment by ID godoc
//
//	@Summary		Get Shipment by a given ID
//	@Description	Get a shipment by its ID
//	@Tags			shipment
//	@Produce		json
//	@Success		200	{object}	Shipment
//	@Router			/shipments/{id} [get]
func GetShipment(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Get Shipment Status by ID godoc
//
//	@Summary		Get Shipment Status by a given ID
//	@Description	Get the status of a shipment by its ID
//	@Tags			shipment
//	@Produce		json
//	@Success		200	{object}	Shipment
//	@Router			/shipments/{id}/status [get]
func GetShipmentStatus(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Update Shipment by ID godoc
//
//	@Summary		Update Shipment by a given ID
//	@Description	Update a shipment by its ID
//	@Tags			shipment
//	@Produce		json
//	@Success		200	{object}	Shipment
//	@Router			/shipments/:id [put]
func UpdateShipment(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Delete Shipment by ID godoc
//
//	@Summary		Delete Shipment by a given ID
//	@Description	Delete a shipment by its ID
//	@Tags			shipment
//	@Produce		json
//	@Success		202
//	@Router			/shipments/:id [delete]
func DeleteShipment(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}
