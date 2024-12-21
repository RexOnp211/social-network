package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
)

type LoginResponse struct {
	Id int `json:"id"`
    Username string `json:"username"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user credentials from the request body
	// Validate the user credentials
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing multipart form:", err)
		return
	}

	username := r.FormValue("email")
	password := r.FormValue("password")
	fmt.Println("TEST", username, password)

	user, err := db.LoginUserDB(username, password)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		log.Printf("Login failed: %v", err)
		return
	}
	fmt.Println("TEST", user)
/* 	userID, err := db.GetUserIDByUsernameOrEmail(username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get user ID"))
		return
	} */
	token, err := NewSession(w, user.Username, user.UserId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create session"))
		return
	}

	log.Printf("User %s logged in with session token %s", user.Username, token)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

    response := LoginResponse{
		Id: user.UserId,
        Username: user.Username,
    }

    w.Header().Set("Content-Type", "application/json")
    if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "Failed to encode response", http.StatusInternalServerError)
        return
    }
}
