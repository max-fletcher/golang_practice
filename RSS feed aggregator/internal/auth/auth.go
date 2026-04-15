package auth

import (
	"fmt"
	"net/http"
	"strings"
)

// extracts api_key from the header of an HTTP request
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")

	if authHeader == "" {
		return "", fmt.Errorf("You are not authenticated")
	}

	vals := strings.Split(authHeader, " ")

	if len(vals) != 2 {
		return "", fmt.Errorf("Malformed auth header")
	}
	if vals[0] != "Bearer" && vals[0] != "bearer" {
		return "", fmt.Errorf("Malformed auth header")
	}

	return vals[1], nil
}
