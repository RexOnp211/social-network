package api

import (
	"net/http"
	"social-network/internal/api/handlers"
)

func Router() {
	http.Handle("/", corsHandler(http.HandlerFunc(handlers.HomeHandler)))
	http.Handle("/posts", corsHandler(http.HandlerFunc(handlers.HandlePosts)))
}

func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Allow-Methods", "Get, Post, Put, Delete")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
