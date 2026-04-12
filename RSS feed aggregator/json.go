package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// This function will take a message(string) and malshal it into a structured JSON string that will be sent as bytes of data
func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with a 5XX error:", msg)
	}

	respondWithJSON(w, code, response{
		Code:    code,
		Status:  "error",
		Message: msg,
	})
}

// This function will take a payload(struct) and malshal it into a JSON string that will be sent as bytes of data
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json response: %v", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
