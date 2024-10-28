package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type CredentialResponse struct {
	Username  string   `json:"username"`
}

func GetCredential(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching credentials...")

	username := ValidateSession(w, r)

	response := CredentialResponse{
		Username:  username,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}