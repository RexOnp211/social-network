package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	db "social-network/pkg/db/sqlite"
)

func GetNicknameFromId(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	fmt.Println("URL", url)
	re := regexp.MustCompile(`\d+$`)
	userId := re.FindString(url)
	nickname := db.GetNicknameFromId(userId)
	fmt.Println("username", userId)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(nickname))
}
