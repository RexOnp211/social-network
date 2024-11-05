package handlers

import (
	"fmt"
	"net/http"

	db "social-network/pkg/db/sqlite"
)

func PrivacyHandler(w http.ResponseWriter, r *http.Request) {

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

	username := r.FormValue("username")
	privacyStatus := r.FormValue("privacy")
	fmt.Println("TEST", username, privacyStatus)

	db.UpdateUserPrivacy(username, privacyStatus)
	
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Privacy setting updated successfully"))
}