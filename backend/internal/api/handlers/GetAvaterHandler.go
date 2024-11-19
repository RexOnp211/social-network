package handlers

import (
	"fmt"
	"net/http"
	"regexp"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"
)

func GetAvaterFromUserId(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	re := regexp.MustCompile(`\d+$`)
	userId := re.FindString(url)
	avatar, err := db.GetAvatarFromUserId(userId)
	if err != nil {
		fmt.Println("Error getting avatar", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pathToFile := "../../assets/image/avatar/" + avatar
	if avatar == "" {
		pathToFile = "../../assets/image/default/profile-default.png"
	}
	buf, err := helpers.EncodeImg(w, pathToFile)
	if err != nil {
		fmt.Println("Error encoding avatar", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error encoding avatar", err)
	}
	w.Write(buf.Bytes())
}
