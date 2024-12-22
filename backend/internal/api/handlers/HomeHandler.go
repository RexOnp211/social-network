package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"
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
	fmt.Println("validate session broken?", nickname)
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

	fmt.Println(filepath)
}
