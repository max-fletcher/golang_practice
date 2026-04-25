package main

import (
	"database/sql"
	"dummy/formatters"
	"dummy/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// handler method that handles feed creation. The addition of (apiCfg apiCfg)
// turns it into a method for apiConfig.
func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
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
	feedFollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow: %v", err))
		return
	}

	respondWithJSON(w, 201, successResponse{
		Code:   201,
		Status: "ok",
		Data:   formatters.DatabaseFeedFollowToFeedFollow(feedFollow),
	})
}

// handler method that handles fetching all feed follows based on user_id. The addition of (apiCfg apiCfg)
// turns it into a method for apiConfig.
func (apiCfg *apiConfig) handlerGetFeedFollowsByUserID(w http.ResponseWriter, r *http.Request, user database.User) {
	// 1st param: context for the request
	feedFollows, err := apiCfg.DB.GetFeedFollowsByUserId(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to fetch feed follows: %v", err))
		return
	}

	respondWithJSON(w, 200, successResponse{
		Code:   200,
		Status: "ok",
		Data:   formatters.DatabaseFeedFollowsToFeedFollows(feedFollows),
	})
}

func (apiCfg *apiConfig) handleDeleteFeedFollowsById(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIdString := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFollowIdString)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing feed follow ID: %v", err))
		return
	}

	// The error handling below is how the deletion logic should be handled i.e if delete command returns a zero, it should mention that it doesn't exist.
	// Also, this(if err == sql.ErrNoRows{...}) is how sqlc checks if no rows were returned as a result of a query
	err = apiCfg.DB.DeleteFeedFollowsById(r.Context(),
		database.DeleteFeedFollowsByIdParams{
			ID:     feedFollowId,
			UserID: user.ID,
		})
	if err != nil {
		if err == sql.ErrNoRows {
			respondWithError(w, 404, fmt.Sprintf("Feed follow with ID %v not found", feedFollowId))
			return
		}
		respondWithError(w, 400, fmt.Sprintf("Failed to delete feed follow with ID %v: %v", feedFollowId, err))
		return
	}

	respondWithJSON(w, 200, successResponse{
		Code:   200,
		Status: "ok",
		Data:   fmt.Sprintf("Feed follow with ID %v deleted sucessfully", feedFollowId),
	})
}
