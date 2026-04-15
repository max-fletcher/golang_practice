package auth

import (
	"context"
	"dummy/internal/auth"
	"dummy/internal/database"
	"fmt"
	"net/http"
)

func AuthenticatedMiddleware(db *database.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			apiKey, err := auth.GetAPIKey(r.Header)
			if err != nil {
				http.Error(w, fmt.Sprintf("Auth error: %v", err), http.StatusForbidden)
				return
			}

			user, err := db.GetUserByAPIKey(r.Context(), apiKey)
			if err != nil {
				http.Error(w, fmt.Sprintf("User not found: %v", err), http.StatusUnauthorized)
				return
			}

			// optionally store user in context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user", user)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
