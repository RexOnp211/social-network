package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
)

func Notifications(w http.ResponseWriter, r *http.Request) {
	fmt.Println("validating session in notifications")
	username := ValidateSession(w, r)
	fmt.Println("userbname", username)
	user, err := db.GetUserFromDb(username)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("failed to get user info from db", err)
		return
	}
	fmt.Println("user", user)
	followRequests, err := db.GetFollowRequestsFromDb(user.Id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Println("failed to get follow requests from db", err)
		return
	}
	fmt.Println("FollowRequests", followRequests)
	json.NewEncoder(w).Encode(followRequests)
}
