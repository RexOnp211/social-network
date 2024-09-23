package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"social-network/pkg"
	db "social-network/pkg/db/sqlite"

	"github.com/gofrs/uuid"
)

// gets posts from the database for the home page
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	Posts := []pkg.Post{
		{
			Id:       1,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       2,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       3,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       4,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       5,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       6,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       7,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       8,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       9,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       10,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       11,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       12,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       13,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       14,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       15,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       16,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       17,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       18,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       19,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       20,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       21,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       22,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       23,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
		{
			Id:       24,
			Title:    "Hello World",
			PostBody: "This is a test post",
			Image:    "https://via.placeholder.com/150",
			Privacy:  "public",
		},
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(Posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing form", err)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error getting file", err)
		return
	}
	defer file.Close()
	subject := r.FormValue("postTitle")
	content := r.FormValue("postBody")
	privacy := r.FormValue("privacy")

	filepath, err := pkg.SaveFile(file, header, "post")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error saving file", err)
		return
	}

	post_id, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error generating uuid", err)
		return
	}

	userId := 1

	data := []interface{}{post_id, userId, subject, content, privacy, filepath}
	err = db.AddPostToDb(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error adding post to db", err)
		return
	}

	fmt.Println(filepath)
}
