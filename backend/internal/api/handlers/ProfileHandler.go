package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"social-network/pkg"
	db "social-network/pkg/db/sqlite"
	"strings"
)

type UserResponse struct {
	User  pkg.User   `json:"user"`
	Posts []pkg.Post `json:"posts"`
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