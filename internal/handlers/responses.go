package handlers

import (
	"encoding/json"
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, body interface{}) {
	jsonResponse, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func respondWithBadRequest(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusBadRequest)
}

func respondWithInternalServerError(w http.ResponseWriter, msg string) {
	http.Error(w, msg, http.StatusInternalServerError)
}
