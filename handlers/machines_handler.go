package handlers

import (
	"encoding/json"
	"log"
	"main/models"
	"main/services"
	"net/http"
)

func HandleMachineData(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received %s request from %s to %s", r.Method, r.RemoteAddr, r.URL.Path)

	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method is allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Body == nil || r.ContentLength == 0 {
		http.Error(w, "Request body cannot be empty", http.StatusBadRequest)
		return
	}

	var machines []models.Machine
	if err := json.NewDecoder(r.Body).Decode(&machines); err != nil {
		http.Error(w, "Error decoding request body: "+err.Error(), http.StatusBadRequest)
		return
	}

	outliers, err := services.DetectOutliers(machines)
	if err != nil {
		http.Error(w, "Error detecting outliers: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(outliers)
}
