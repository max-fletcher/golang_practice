package main

import (
	"net/http"
)

// This function will take a payload(struct) and malshal it into a JSON string that will be sent as bytes of data
func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	type response struct {
		// We are using `json:"status"` because without this, we will get { "Status": "ok" }
		Code    int    `json:"code"`
		Status  string `json:"status"`
		Message string `json:"message"`
	}

	respondWithJSON(w, 200, response{
		Code:    200,
		Status:  "ok",
		Message: "Server running",
	})
}
