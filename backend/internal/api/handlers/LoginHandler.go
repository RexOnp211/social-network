package handlers

import (
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"time"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Get the user credentials from the request body
	// Validate the user credentials
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	login, err := db.LoginUserDB(username, password)
	if err != nil {
		fmt.Println("Error logging in", err)
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(err.Error()))
		log.Printf("Login failed: %v", err)
		return
	}
	userID, err := db.GetUserIDByUsernameOrEmail(username)
	if err != nil {
		fmt.Println("Error getting user ID", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to get user ID"))
		return
	}
	token, err := NewSession(w, login.Username, userID)
	if err != nil {
		fmt.Println("Error creating session", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to create session"))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    token,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	})

	log.Printf("User %s logged in with session token %s", login.Username, token)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(token))

}
