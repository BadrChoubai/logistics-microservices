package handler

import (
	"encoding/json"
	"net/http"
)

type HealthcheckResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

type Item struct{}

// Inventory Service Health godoc
//
//	@Summary		Health check
//	@Description	Returns service health check response
//	@Tags			system
//	@Produce		json
//	@Success		200	{object}	HealthcheckResponse
//	@Router			/health [get]
//	@Router			/health [head]
func GetInventoryHealth(w http.ResponseWriter, r *http.Request) {
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
//	@Summary		Get Items
//	@Description	Returns a list of all items
//	@Tags			inventory
//	@Produce		json
//	@Success		200	{array}	Item
//	@Router			/items [get]
func GetItems(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Create Item godoc
//
//	@Summary		Create Item
//	@Description	Create a new item
//	@Tags			inventory
//	@Produce		json
//	@Success		201	{object}	Item
//	@Router			/items [post]
func CreateItem(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Get Item godoc
//
//	@Summary		Get Item
//	@Description	Get an item from inventory using its id
//	@Tags			inventory
//	@Produce		json
//	@Success		200	{object}	Item
//	@Router			/items/{id} [get]
func GetItem(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not implemented", http.StatusNotImplemented)
}

// Update Item godoc
//
//	@Summary		update item
//	@Description	update an item from inventory using its id
//	@Tags			inventory
//	@Produce		json
//	@Success		204	{object}	Item
//	@Router			/items/{id} [put]
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

// Delete Item godoc
//
//	@Summary		delete item
//	@Description	delete an item from inventory using its id
//	@Tags			inventory
//	@Produce		json
//	@Success		204	{object}	Item
//	@Router			/items/{id} [delete]
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

// Update Item Stock godoc
//
//	@Summary		update item stock
//	@Description	update an item's stock count from inventory using its id
//	@Tags			inventory
//	@Produce		json
//	@Success		204	{object}	Item
//	@Router			/items/{id}/stock [put]
func UpdateItemStock(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}

// Get Low Stock Items godoc
//
//	@Summary		list items from inventory with low stock
//	@Description	List items from inventory with low stock count
//	@Tags			inventory
//	@Produce		json
//	@Success		200	{array}	Item
//	@Router			/items/low-stock [get]
func GetLowStockItems(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "not implemented", http.StatusNotImplemented)
}
