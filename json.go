package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error while marshalling the response", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
}

func respondError(w http.ResponseWriter, status int, message string) {
	if status > 499 {
		log.Println("Responding with 5XX error: ", message)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondJSON(w, status, errorResponse{
		Error: message,
	})
}
