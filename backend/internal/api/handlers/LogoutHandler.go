package handlers

import (
	"fmt"
	"net/http"
)

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	CloseSession(w, r)
	fmt.Println("User logged out")
}
