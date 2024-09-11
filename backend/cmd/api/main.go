package main

import (
	"fmt"
	"net/http"
	"social-network/internal/api"
)

func main() {
	api.Router()

	fmt.Println("starting go server")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
