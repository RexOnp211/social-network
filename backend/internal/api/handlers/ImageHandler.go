package handlers

import (
	"bytes"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

// to use this handler send get request to /image/{imagename}
// TODO: add functionality to get avatar images
func GetImageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("getting file", r.URL.Path[7:])
	pathToFile := "../../assets/image/upload/" + r.URL.Path[7:]
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println("Error opening file", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var buf bytes.Buffer
	switch pathToFile[len(pathToFile)-3:] {
	case "png":
		img, err := png.Decode(file)
		if err != nil {
			fmt.Println("Error decoding png", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = png.Encode(&buf, img)
		if err != nil {
			fmt.Println("Error encoding png", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "image/png")
	case "jpg", "peg":
		img, err := jpeg.Decode(file)
		if err != nil {
			fmt.Println("Error decoding jpeg", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = jpeg.Encode(&buf, img, nil)
		if err != nil {
			fmt.Println("Error encoding jpeg", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "image/jpeg")
	case "gif":
		gifImg, err := gif.DecodeAll(file)
		if err != nil {
			fmt.Println("Error decoding gif", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = gif.EncodeAll(&buf, gifImg)
		if err != nil {
			fmt.Println("Error encoding gif", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "image/gif")
	default:
		fmt.Println("File type not supported")
		http.Error(w, "File type not supported", http.StatusUnsupportedMediaType)
		return
	}

	w.Write(buf.Bytes())
}
