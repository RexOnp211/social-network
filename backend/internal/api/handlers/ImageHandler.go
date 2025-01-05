package handlers

import (
	"fmt"
	"net/http"
	"social-network/pkg/helpers"
)

// to use this handler send get request to /image/{imagename}
// TODO: add functionality to get avatar images
func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting file", r.URL.Path[7:])
	pathToFile := "../../assets/image/upload/" + r.URL.Path[7:]
	buf, err := helpers.EncodeImg(w, pathToFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error encoding image", err)
		return
	}
	w.Write(buf.Bytes())
}

func GetGroupImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting file", r.URL.Path[7:])
	pathToFile := "../../assets/image/upload/group-post-image/" + r.URL.Path[17:]
	buf, err := helpers.EncodeImg(w, pathToFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error encoding image", err)
		return
	}
	w.Write(buf.Bytes())
}


func GetGroupCommentImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting file", r.URL.Path[7:])
	pathToFile := "../../assets/image/upload/group-post-comment-image/" + r.URL.Path[25:]
	buf, err := helpers.EncodeImg(w, pathToFile)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("Error encoding image", err)
		return
	}
	w.Write(buf.Bytes())
}