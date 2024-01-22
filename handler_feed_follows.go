package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/LE0-MAhendra/gotestproj/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}
	feedfollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn`t create Feed Follow: %v", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollow(feedfollow))
}

func (apiCfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn`t get Feed Follows: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedFollows(feedFollows))
}
func (apiCfg *apiConfig) handlerDeleteFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFloowstr := chi.URLParam(r, "feedFollowId")
	feedFollowId, err := uuid.Parse(feedFloowstr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn`t parse Feed Follows Id : %v", err))
		return
	}
	err = apiCfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     feedFollowId,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn`t delete Feed Follow : %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})
}
