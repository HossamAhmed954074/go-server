package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responcedwithJson(w http.ResponseWriter, status int, data interface{}) {
	dataJson, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error converting data to JSON: %v", err)
		http.Error(w, "Error converting data to JSON", http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(dataJson)
}

func responcedwithError(w http.ResponseWriter, status int, message string) {
	type errorResponse struct {
		Error string `json:"error"`
	}
	responcedwithJson(w, status, errorResponse{Error: message})
}
