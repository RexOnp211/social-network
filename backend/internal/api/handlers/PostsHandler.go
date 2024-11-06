package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"
	"strconv"
)

func HandlePosts(w http.ResponseWriter, r *http.Request) {
	strid := r.URL.Path[len("/post/"):]
	id, err := strconv.Atoi(strid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing id", err)
		return
	}
	post, err := db.GetPostFromId(id)
	if err != nil {
		fmt.Println("Error getting post", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	comments, err := db.GetCommentsFromPostId(id)
	if err != nil {
		fmt.Println("Error getting comments", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	post.Comments = comments
	fmt.Println("this is a post", post)
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
	// postUUID, err := uuid.FromString(postId)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// 	fmt.Println("Error parsing uuid", err)
	// 	return
	// }
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

	data := []interface{}{postId, "3213", content, filepath}
	fmt.Println("data before save", data)
	err = db.AddCommentToDb(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error adding comment", err)
		return
	}
}
