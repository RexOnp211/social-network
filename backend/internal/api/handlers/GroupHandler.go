package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"
	"strconv"
	"strings"
	"time"
)

type GroupResponse struct {
	Group helpers.Group `json:"group"`
}

type GroupsResponse struct {
	Groups []helpers.Group `json:"groups"`
}

type GroupPostsResponse struct {
	Posts []helpers.GroupPost `json:"groupPosts"`
}

type Invitation struct {
	ID        int    `json:"id"`
	GroupName string `json:"groupname"`
	UserName  string `json:"username"`
	Status    string `json:"status"`
}

type MemberStatusChange struct {
	GroupName string `json:"groupname"`
	UserName  string `json:"username"`
	Status    string `json:"status"`
}

// BASIC GROUP INFO ----------------------------------

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
		Group: group,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GroupsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching groups...")

	groups, err := db.GetGroupsFromDb()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting groups", err)
		return
	}

	response := GroupsResponse{
		Groups: groups,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateGroupHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(50 << 20) // 50MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing multipart form:", err)
		return
	}

	creator := r.FormValue("user")
	title := r.FormValue("title")
	description := r.FormValue("description")
	fmt.Println("TEST", creator, title, description)

	data := []interface{}{creator, title, description}
	err = db.CreateGroupDB(data)
	if err != nil {
		log.Printf("Create group failed: %v", err)
		msg := "Failed to create group. Please retry later."
		if err.Error() == "UNIQUE constraint failed: groups.title" {
			msg = fmt.Sprintf(`Group named "%v" already exists.`, title)
		}
		http.Error(w, msg, http.StatusInternalServerError)
		return
	}

	user, err := db.GetUserFromDb(creator)
	if err != nil {
		fmt.Println("Error Getting user from db")
		return
	}
	chatId, err := db.GetGroupIdWithCreatorName(creator, title)
	if err != nil {
		fmt.Println("Error getting chatid from db")
		return
	}

	err = db.AddUserIntoChatRoom(user.Id, chatId)
	if err != nil {
		fmt.Println("error adding chat creator into chat", err)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("ERROR getting user from db", err)
	}

	log.Printf("User %s created group %s", creator, title)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}

// MEMBERSHIPS ----------------------------------

func MembershipsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching group members...")

	nickname := strings.TrimPrefix(r.URL.Path, "/fetch_memberships/")
	if nickname == "" {
		http.Error(w, "Username not provided", http.StatusBadRequest)
		return
	}

	memberships, err := db.GetMembershipsFromDb(nickname)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting membership", err)
		return
	}

	response := struct {
		Memberships []helpers.Membership `json:"memberships"`
	}{
		Memberships: memberships,
	}

	log.Println("response from fetch memberships", response)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		fmt.Println("Error encoding JSON", err)
	}
}

func InviteMemberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	invitation := struct {
		Groupname string `json:"groupname"`
		Username  string `json:"username"`
	}{}
	err := json.NewDecoder(r.Body).Decode(&invitation)
	if err != nil {
		log.Println("Failed to decode JSON")
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	log.Println("inviting member...", invitation)
	chatId, err := db.GetChatIdFromGroup(invitation.Groupname)
	if err != nil {
		fmt.Println("Error getting chatid from group", err)
		return
	}
	errorMessage, isError := db.InviteMemberDB(invitation.Groupname, invitation.Username, chatId)
	fmt.Println("TEST", errorMessage, isError)
	if isError {
		if errorMessage != "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{"message": errorMessage}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		response := map[string]string{"message": "Internal server error"}
		json.NewEncoder(w).Encode(response)
		log.Printf("Invite member failed: %v", errorMessage)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": fmt.Sprintf(`User "%s" invited to group %s`, invitation.Username, invitation.Groupname)}
	json.NewEncoder(w).Encode(response)
}

func UpdateMemberStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("changing member stat... 1")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	log.Println("changing member stat... 2")

	data := struct {
		ID        int
		Groupname string
		Username  string
		Status    string
		ChatId    int
	}{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		log.Println("Failed to decode JSON")
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	log.Println("changing member stat... 3", data)

	db.UpdateMemberStatus(data.ID, data.Groupname, data.Username, data.Status, data.ChatId)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Member status changed successfully"))
}

// POST & COMMENT ----------------------------------

func CreateGroupPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("posting to group...")

	err := r.ParseMultipartForm(50 << 20) // 50 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing form", err)
		return
	}

	userID := r.FormValue("user_id")
	groupname := r.FormValue("groupname")
	nickname := r.FormValue("nickname")
	subject := r.FormValue("postTitle")
	content := r.FormValue("postBody")

	file, header, err := r.FormFile("image")
	log.Println("TEST", file, header, err)

	filepath := ""
	if err == nil {
		defer file.Close()
		filepath, err = helpers.SaveFile(file, header, "group-post")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Error saving file", err)
			return
		}
	}
	data := []interface{}{groupname, userID, nickname, subject, content, filepath}
	log.Println("after savefile")
	err = db.AddGroupPostToDb(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error adding group post to db", err)
		return
	}

	fmt.Println(filepath)
}

func GroupPostsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching group posts...", r.URL.Path)

	path := strings.TrimPrefix(r.URL.Path, "/fetch_group_posts/")
	if path == "" {
		http.Error(w, "Groupname not provided", http.StatusBadRequest)
		return
	}
	log.Println(path)

	posts, err := db.GetGroupPostsFromDb(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting groups 1", err)
		return
	}

	response := GroupPostsResponse{
		Posts: posts,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GroupPostHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching group post...", r.URL.Path)

	path := strings.TrimPrefix(r.URL.Path, "/fetch_group_post/")
	if path == "" {
		http.Error(w, "Post ID not provided", http.StatusBadRequest)
		return
	}
	log.Println(path)
	intID, _ := strconv.Atoi(path)

	post, err := db.GetGroupPostFromDb(intID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting post", err)
		return
	}

	response := struct {
		Post helpers.GroupPost `json:"post"`
	}{
		Post: post,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GroupPostCommentsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching group comments...", r.URL.Path)

	path := strings.TrimPrefix(r.URL.Path, "/fetch_group_post_comment/")
	if path == "" {
		http.Error(w, "Post ID not provided", http.StatusBadRequest)
		return
	}
	log.Println(path)
	intID, _ := strconv.Atoi(path)

	comments, err := db.GetGroupPostCommentsFromDb(intID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting comments", err)
		return
	}

	response := struct {
		Comments []helpers.GroupComment `json:"comments"`
	}{
		Comments: comments,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateGroupPostCommentHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("posting comment...")

	err := r.ParseMultipartForm(50 << 20) // 50 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing form", err)
		return
	}

	postID := r.FormValue("post_id")
	userID := r.FormValue("user_id")
	nickname := r.FormValue("nickname")
	content := r.FormValue("commentBody")

	file, header, err := r.FormFile("image")

	filepath := ""
	if err == nil {
		defer file.Close()
		filepath, err = helpers.SaveFile(file, header, "group-post-comment")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Error saving file", err)
			return
		}
	}
	data := []interface{}{postID, userID, nickname, content, filepath}
	log.Println("after savefile")
	err = db.AddGroupPostCommentToDb(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error adding comment to db", err)
		return
	}

	fmt.Println(filepath)
}

// EVENTS ----------------------------------

func FetchGroupEventsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching:", r.URL.Path)

	target := strings.TrimPrefix(r.URL.Path, "/fetch_group_events/")
	if target == "" {
		http.Error(w, "target not provided", http.StatusBadRequest)
		return
	}
	log.Println(target)

	events := db.GetGroupEventsFromDb(target)

	response := struct {
		Data []helpers.GroupEvent `json:"data"`
	}{
		Data: events,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CreateGroupEventHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("creating events...")

	err := r.ParseMultipartForm(50 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing form", err)
		return
	}

	groupname := r.FormValue("groupname")
	userID := r.FormValue("user_id")
	nickname := r.FormValue("nickname")
	title := r.FormValue("eventTitle")
	description := r.FormValue("eventDescription")
	date := r.FormValue("eventDate")

	// Validate eventDate
	eventTime, err := time.Parse("2006-01-02T15:04", date)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		log.Println("Error parsing date:", err)
		return
	}

	// Check if the date is in the past
	if eventTime.Before(time.Now()) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		response := map[string]string{"message": "Event date cannot be in the past"}
		json.NewEncoder(w).Encode(response)
		return
	}

	data := []interface{}{groupname, userID, nickname, title, description, date}
	err = db.AddGroupEventToDb(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error adding group event to db", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Event created successfully"))
}

func FetchUserEventStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching event status...")

	nickname := r.URL.Query().Get("username")
	event := r.URL.Query().Get("event")

	if nickname == "" || event == "" {
		http.Error(w, "Missing parameters", http.StatusBadRequest)
		return
	}
	eventID, _ := strconv.Atoi(event)
	log.Println(nickname, eventID)

	eventStatus := db.GetEventStatusFromDb(nickname, eventID)
	log.Println(eventStatus)

	log.Println("response", eventStatus)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(eventStatus); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		fmt.Println("Error encoding JSON", err)
	}
}

func UpdateEventStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("changing event stat...")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var eventStatus helpers.EventStatus
	err := json.NewDecoder(r.Body).Decode(&eventStatus)
	if err != nil {
		log.Println("Failed to decode JSON")
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	log.Println("changing event stat...", eventStatus)

	db.UpdateEventStatus(eventStatus.Nickname, eventStatus.EventId, eventStatus.Going)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Member status changed successfully"))
}

func FetchYourRequestsHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching group you made...")

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var groups []helpers.Membership
	err := json.NewDecoder(r.Body).Decode(&groups)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	log.Println("Received groups:", groups)

	groupsYouMade := db.GetYourRequestsFromDb(groups)
	log.Println(groupsYouMade)

	log.Println("response", groupsYouMade)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(groupsYouMade); err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		fmt.Println("Error encoding JSON", err)
	}
}
