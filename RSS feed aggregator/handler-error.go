package main

import (
	"net/http"
)

// This function will return an error response no matter what.
func handlerError(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 500, "Something went wrong.")
}
