package handler

import (
	"encoding/json"
	"net/http"
)

type HealthcheckResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// Health godoc
//
//	@Summary		Health check
//	@Description	Returns service health check response
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	HealthcheckResponse
//	@Router			/health [get]
//	@Router			/health [head]
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodHead {
		w.WriteHeader(http.StatusOK)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	// nolint
	json.NewEncoder(w).Encode(HealthcheckResponse{Status: "Healthy", Version: "1.0.0"})
}
