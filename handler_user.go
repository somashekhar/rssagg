package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/somashekhar/rssagg/internal/auth"
	"github.com/somashekhar/rssagg/internal/database"
)

func (apiConfig *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`
	}
	params := parameters{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON:%v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      params.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseUserToUser(user))
}

func (apiConfig *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusForbidden, fmt.Sprintf("Error getting API key: %v", err))
		return
	}

	user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, fmt.Sprintf("Error getting user: %v", err))
		return
	}

	respondWithJSON(w, http.StatusOK, databaseUserToUser(user))
}
