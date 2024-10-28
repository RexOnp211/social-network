package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"

	"github.com/gofrs/uuid"
)

func HandlePosts(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/post/"):]
	uuid, err := uuid.FromString(id)
	if err != nil {
		fmt.Println("Error parsing uuid", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post, err := db.GetPostFromId(uuid)
	if err != nil {
		fmt.Println("Error getting post", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(post)
	comments, err := db.GetCommentsFromPostId(uuid)
	if err != nil {
		fmt.Println("Error getting comments", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(comments)
	json.NewEncoder(w).Encode(post)
}

func CreateComment(w http.ResponseWriter, r *http.Request) {
	fmt.Println("this shit works")
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing form", err)
		return
	}
	postId := r.URL.Path[len("/post/"):]
	postUUID, err := uuid.FromString(postId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing uuid", err)
		return
	}
	content := r.FormValue("commentBody")
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
	comment_id, err := uuid.NewV4()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error generating uuid", err)
		return
	}

	userId := 1

	data := []interface{}{comment_id, postUUID, userId, content, filepath}
	fmt.Println("data before save", data)
	err = db.AddCommentToDb(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error adding comment", err)
		return
	}
}
