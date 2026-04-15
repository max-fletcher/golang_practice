package main

import (
	"dummy/formatters"
	"dummy/internal/auth"
	"dummy/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// handler method that handles user creation. The addition of (apiCfg apiCfg)
// turns it into a method for apiConfig.
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body) // decode request body
	params := parameters{}
	err := decoder.Decode(&params) // decode request body(json) and store in params variable, or get an error
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	// 1st param: context for the request
	// 2nd param: the struct that we want to pass so it saves the underlying data in DB
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}

	respondWithJSON(w, 201, successResponse{
		Code:   201,
		Status: "ok",
		Data:   formatters.DatabaseUserToUser(user),
	})
}

// handler method that handles user creation. The addition of (apiCfg apiCfg) turns it into a method for apiConfig.
// This is now used to show another way to implement middleware(see route with "handlerGetUserByAPIKey2" func)
func (apiCfg *apiConfig) handlerGetUserByAPIKey2(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	// 1st param: context for the request
	// 2nd param: the struct that we want to pass so it saves the underlying data in DB
	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, 401, fmt.Sprintf("User not found: %v", err))
		return
	}

	respondWithJSON(w, 200, successResponse{
		Code:   200,
		Status: "ok",
		Data:   formatters.DatabaseUserToUser(user),
	})
}

// Turned this to a different function that canwill only be used in an auth middleware(see auth.go in root folder and "handlerGetUserByAPIKey" route)
func (apiCfg *apiConfig) handlerGetUserByAPIKey(w http.ResponseWriter, r *http.Request, user database.User) {
	respondWithJSON(w, 200, successResponse{
		Code:   200,
		Status: "ok",
		Data:   formatters.DatabaseUserToUser(user),
	})
}
