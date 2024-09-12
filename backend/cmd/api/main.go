package main

import (
	"fmt"
	"net/http"
	"social-network/internal/api"
	"social-network/internal/api/handlers"
)

func main() {
	r := &api.Router{}

	r.AddRoute("GET", "/posts", http.HandlerFunc(handlers.HandlePosts))
	r.AddRoute("GET", "/", http.HandlerFunc(handlers.HomeHandler))

	fmt.Println("starting go server")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		fmt.Println(err)
	}
}
