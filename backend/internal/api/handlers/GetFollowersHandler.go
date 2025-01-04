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
	fmt.Println("this is the user array for follower users", userArr)
	json.NewEncoder(w).Encode(userArr)
}
