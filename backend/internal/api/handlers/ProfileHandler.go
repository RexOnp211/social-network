package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"
	"strings"
)

type UserResponse struct {
	User  helpers.User   `json:"user"`
	Posts []helpers.Post `json:"posts"`
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching user...")
	log.Println(r.URL.Path)

	path := strings.TrimPrefix(r.URL.Path, "/profile/")
	if path == "" {
		http.Error(w, "Username not provided", http.StatusBadRequest)
		return
	}

	log.Println(path)

	user, err := db.GetUserFromDb(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting user", err)
		return
	}

	// user not found
	log.Println("user default", user)
	if (user == helpers.User{}) {
		http.Error(w, "No user found with nickname:", http.StatusInternalServerError)
		log.Println("No user found with nickname:", err)
		return
	}

	// fetch posts by user_id
	posts, err := db.GetUserPostFromDbByUser(user.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting user's post", err)
		return
	}

	// TODO: fetch followers
	/* followers, err := db.GetFollowersFromDb(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting followers", err)
		return
	}
	*/
	// TODO: fetch followings
	/* followings, err := db.GetFollowingsFromDb(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting followings", err)
		return
	} */

	response := UserResponse{
		User:  user,
		Posts: posts,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
