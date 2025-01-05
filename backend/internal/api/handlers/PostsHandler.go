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
	nickname := ValidateSession(w, r)
	user, err := db.GetUserFromDb(nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting user", err)
		return
	}
	data := []interface{}{postId, user.Id, content, filepath}
	fmt.Println("data before save", data)
	err = db.AddCommentToDb(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error adding comment", err)
		return
	}
}

func SetPostPrivacy(w http.ResponseWriter, r *http.Request) {
	// Parse request body
	type PrivacyRequest struct {
		PostID         int   `json:"post_id"`
		AllowedUserIDs []int `json:"allowed_user_ids"`
	}

	var req PrivacyRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		log.Println("Error decoding request in SetPostPrivacy:", err)
		return
	}

	// Validate session to ensure the request is coming from the post creator
	nickname := ValidateSession(w, r)
	user, err := db.GetUserFromDb(nickname)
	if err != nil {
		http.Error(w, "Failed to fetch user data", http.StatusInternalServerError)
		log.Println("Error fetching user in SetPostPrivacy:", err)
		return
	}

	// Check if the user is the owner of the post
	post, err := db.GetPostFromId(req.PostID)
	if err != nil {
		http.Error(w, "Failed to fetch post data", http.StatusInternalServerError)
		log.Println("Error fetching post in SetPostPrivacy:", err)
		return
	}
	if post.UserId != strconv.Itoa(user.Id) {
		http.Error(w, "Unauthorized to modify this post", http.StatusUnauthorized)
		log.Println("Unauthorized access in SetPostPrivacy")
		return
	}

	// Save privacy data to the database
	err = db.SavePostPrivacy(req.PostID, req.AllowedUserIDs)
	if err != nil {
		http.Error(w, "Failed to save post privacy settings", http.StatusInternalServerError)
		log.Println("Error saving post privacy in SetPostPrivacy:", err)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Post privacy settings updated successfully"))
}
