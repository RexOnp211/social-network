package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
)

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	username := ValidateSession(w, r)
	user, err := db.GetUserFromDb(username)
	if err != nil {
		log.Println("error getting user Info for following list", err)
		return
	}
	userArr, err := db.GetUsersFollowingListFromDb(user.Id)
	if err != nil {
		log.Println("error getting people user follows", err)
		return
	}
	fmt.Println("this is the user array for following users", userArr)
	json.NewEncoder(w).Encode(userArr)
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	username := ValidateSession(w, r)
	user, err := db.GetUserFromDb(username)
	if err != nil {
		log.Println("error getting user Info for following list", err)
		return
	}
	userArr, err := db.GetUsersFollowersListFromDB(user.Id)
	if err != nil {
		log.Println("error getting people user follows", err)
		return
	}
	fmt.Println("this is the user array for following users", userArr)
	json.NewEncoder(w).Encode(userArr)
}
