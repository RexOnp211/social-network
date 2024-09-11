package handlers

import "net/http"

func HandlePosts(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/posts" {
		return
	}
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Hello forn backend"}`))
	}
}
