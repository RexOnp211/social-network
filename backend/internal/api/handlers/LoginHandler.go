package handlers

import (
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
)

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

	http.SetCookie(w, &http.Cookie{
		Name:  "session_token",
		Value: token,
	})

	log.Printf("User %s logged in with session token %s", user.Username, token)

	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
}
