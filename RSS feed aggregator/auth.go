package main

import (
	"dummy/internal/auth"
	"dummy/internal/database"
	"fmt"
	"net/http"
)

// This line defines a func signature i.e params that a function takes. Dries up the code. The 3rd params is basically
// a user struct so this function is expected to return a user struct.
type authHandler func(w http.ResponseWriter, r *http.Request, userDB database.User)

// Remeber that we are still using main as package(see above) so we are appending this method to the apiConfig struct(i.e making it a method of apiConfig).
// This is how we curry a function. This function takes a handler function as an arg and returns the function(or at least returns the result of that func).
func (apiCfg *apiConfig) authenticatedMiddleware(authHandler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 401, fmt.Sprintf("User not found: %v", err))
			return
		}

		// checking if the user data can be retrieved from context since the request is passing through "AuthenticatedMiddleware"
		userData, _ := r.Context().Value("user").(database.User)
		fmt.Printf("User data from context: %v", userData)

		authHandler(w, r, user)
	}
}
