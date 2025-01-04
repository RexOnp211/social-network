package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"strings"
)

func GetFollowing(w http.ResponseWriter, r *http.Request) {
	nickname := r.URL.Path
	trimmedNickname := strings.TrimPrefix(nickname, "/following/")
	fmt.Println("userId In url", trimmedNickname)
	user, err := db.GetUserFromDb(trimmedNickname)
	if err != nil {
		log.Println("error getting user from db in GetFollowing", err)
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
	nickname := r.URL.Path
	trimmedNickname := strings.TrimPrefix(nickname, "/followers/")
	user, err := db.GetUserFromDb(trimmedNickname)
	if err != nil {
		log.Println("error getting user from db in GetFollowers", err)
		return
	}
	userArr, err := db.GetUsersFollowersListFromDB(user.Id)
	if err != nil {
		log.Println("error getting people user follows", err)
		return
	}
	fmt.Println("this is the user array for followers users", userArr)
	json.NewEncoder(w).Encode(userArr)
}

func GetUnfollowing(w http.ResponseWriter, r *http.Request) {
	var payload struct {
		FollowerID int `json:"follower_id"`
		FolloweeID int `json:"followee_id"`
	}

	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		log.Println("Error decoding unfollow request")
		return
	}

	if payload.FollowerID == 0 || payload.FolloweeID == 0 {
		log.Println("Unfollow request payload bugged out")
		return
	}

	err = db.UnfollowUserFromDB(payload.FollowerID, payload.FolloweeID)
	if err != nil {
		log.Println("Error unfollowing user:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "User unfollowed successfully"}`))
}
