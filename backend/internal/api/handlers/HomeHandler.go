package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"
	"strconv"
)

// gets posts from the database for the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := db.GetPostsFromDb()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting posts", err)
		return
	}
	json.NewEncoder(w).Encode(posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing form", err)
		return
	}

	subject := r.FormValue("postTitle")
	content := r.FormValue("postBody")
	privacy := r.FormValue("privacy")

	file, header, err := r.FormFile("image")
	filepath := ""
	if err == nil {
		defer file.Close()
		filepath, err = helpers.SaveFile(file, header, "post")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Error saving file", err)
			return
		}
	}
	nickname := ValidateSession(w, r)
	user, err := db.GetUserFromDb(nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting user", err)
		return
	}
	data := []interface{}{user.Id, subject, content, filepath, privacy}
	err = db.AddPostToDb(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error adding post to db", err)
		return
	}

	lastInsertID, err := db.GetLastInsertID()
	if err != nil {
		http.Error(w, "Error fetching last post ID", http.StatusInternalServerError)
		log.Println("Error fetching last post ID:", err)
		return
	}

	if privacy == "private" {
		allowedUserIDs := r.Form["allowedUserIDs[]"]
		intUserIDs := []int{}

		for _, userID := range allowedUserIDs {
			id, err := strconv.Atoi(userID)
			if err != nil {
				http.Error(w, "Invalid user ID format", http.StatusBadRequest)
				log.Println("Invalid user ID format:", err)
				return
			}
			intUserIDs = append(intUserIDs, id)
		}

		err = db.SavePostPrivacy(lastInsertID, intUserIDs)
		
	}
	fmt.Println(filepath)
}
