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

type GroupResponse struct {
	Group  pkg.Group   `json:"group"`
}

func GroupHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching group...")
	log.Println(r.URL.Path)

	path := strings.TrimPrefix(r.URL.Path, "/group/")
	if path == "" {
		http.Error(w, "Groupname not provided", http.StatusBadRequest)
		return
	}

	log.Println(path)

	group, err := db.GetGroupFromDb(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting group", err)
		return
	}

	// TODO: fetch posts & event for the group
	/* posts, err := db.GetUserPostFromDbByUser(user.Id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Error getting user's post", err)
			return
		} */

	response := GroupResponse{
		Group:  group,
	}
	log.Println("response", response)


	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}