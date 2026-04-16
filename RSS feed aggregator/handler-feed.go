package main

import (
	"dummy/formatters"
	"dummy/internal/database"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

// handler method that handles feed creation. The addition of (apiCfg apiCfg)
// turns it into a method for apiConfig.
func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
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
	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed: %v", err))
		return
	}

	respondWithJSON(w, 201, successResponse{
		Code:   201,
		Status: "ok",
		Data:   formatters.DatabaseFeedToFeed(feed),
	})
}

// handler method that handles fetching all feeds. The addition of (apiCfg apiCfg)
// turns it into a method for apiConfig.
func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	// 1st param: context for the request
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Failed to fetch feeds: %v", err))
		return
	}

	respondWithJSON(w, 200, successResponse{
		Code:   200,
		Status: "ok",
		Data:   formatters.DatabaseFeedsToFeeds(feeds),
	})
}
