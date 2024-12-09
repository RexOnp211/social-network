package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
)

type CredentialResponse struct {
	Username string `json:"username"`
	Id       int    `json:"id"`
}

func GetCredential(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching credentials...")

	username := ValidateSession(w, r)
	user, err := db.GetUserFromDb(username)
	if err != nil {
		log.Println("Error getting user", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := CredentialResponse{
		Username: username,
		Id:       user.Id,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
