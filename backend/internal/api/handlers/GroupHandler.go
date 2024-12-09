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

type GroupResponse struct {
	Group helpers.Group `json:"group"`
}

type GroupsResponse struct {
	Groups []helpers.Group `json:"groups"`
}

type GroupsMembersResponse struct {
	GroupMembers []helpers.GroupMembers `json:"groupmembers"`
}

type Invitation struct {
	ID     int    `json:"id"`
	Status string `json:"status"`
}

type MemberStatusChange struct {
	GroupName     string    `json:"groupname"`
	UserName string `json:"username"`
	Status string `json:"status"`
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

	err := r.ParseMultipartForm(10 << 20) // 10MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing multipart form:", err)
		return
	}

	creator := r.FormValue("user_id")
	title := r.FormValue("title")
	description := r.FormValue("description")
	fmt.Println("TEST", title, description)

	data := []interface{}{creator, title, description}
	err = db.CreateGroupDB(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		log.Printf("Create group failed: %v", err)
		return
	}

	log.Printf("User %s created group %s", creator, title)
}

func InviteMemberHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var memberStatusChange MemberStatusChange
	err := json.NewDecoder(r.Body).Decode(&memberStatusChange)
	if err != nil {
		log.Println("Failed to decode JSON")
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	log.Println("changing member stat...")

	errorMessage, err := db.InviteMemberDB(memberStatusChange.GroupName, memberStatusChange.UserName, memberStatusChange.Status)
	fmt.Println("TEST",errorMessage, err )
	if err != nil {
		if errorMessage != "" {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{"message": errorMessage}
			json.NewEncoder(w).Encode(response)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			response := map[string]string{"message": "Internal server error"}
			json.NewEncoder(w).Encode(response)
		}
		log.Printf("Invite member failed: %v", err)
		return
	}

	log.Printf("User %s %s to group %s", memberStatusChange.UserName, memberStatusChange.Status, memberStatusChange.GroupName)
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": fmt.Sprintf("User %s %s to group %s", memberStatusChange.UserName, memberStatusChange.Status, memberStatusChange.GroupName)}
	json.NewEncoder(w).Encode(response)
}

func GroupMemberHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("fetching group members...")

	path := strings.TrimPrefix(r.URL.Path, "/group_member/")
	if path == "" {
		http.Error(w, "Username not provided", http.StatusBadRequest)
		return
	}

	invitations, err := db.GetGroupMembersFromDb(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error getting group members", err)
		return
	}

	response := GroupsMembersResponse{
		GroupMembers: invitations,
	}
	log.Println("response", response)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func UpdateMemberStatusHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("changing member stat...")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	log.Println("changing member stat...")


	var invitation Invitation
	err := json.NewDecoder(r.Body).Decode(&invitation)
	if err != nil {
		log.Println("Failed to decode JSON")
		http.Error(w, "Failed to decode JSON", http.StatusBadRequest)
		return
	}
	log.Println("changing member stat...")


	err = db.UpdateMemberStatus(invitation.ID, invitation.Status)
	if err != nil {
		http.Error(w, "Failed to update status", http.StatusInternalServerError)
		return
	}
	log.Println("changing member stat...")


	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Member status changed successfully"))
}