package helpers

import (
	"bytes"
	"fmt"
	"image/gif"
	"image/jpeg"
	"image/png"
	"net/http"
	"os"
)

func EncodeImg(w http.ResponseWriter, pathToFile string) (bytes.Buffer, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		fmt.Println("Error opening file", err)
		return bytes.Buffer{}, err
	}
	defer file.Close()

	var buf bytes.Buffer
	fmt.Println("pathToFil2e", pathToFile)
	fmt.Println("pathToFile", pathToFile[len(pathToFile)-3:])
	switch pathToFile[len(pathToFile)-3:] {
	case "png":
		img, err := png.Decode(file)
		if err != nil {
			fmt.Println("Error decoding png", err)
			return bytes.Buffer{}, err
		}
		err = png.Encode(&buf, img)
		if err != nil {
			fmt.Println("Error encoding png", err)
			return bytes.Buffer{}, err
		}
		w.Header().Set("Content-Type", "image/png")
	case "jpg", "peg":
		img, err := jpeg.Decode(file)
		if err != nil {
			fmt.Println("Error decoding jpeg", err)
			return bytes.Buffer{}, err
		}
		err = jpeg.Encode(&buf, img, nil)
		if err != nil {
			fmt.Println("Error encoding jpeg", err)
			return bytes.Buffer{}, err
		}
		w.Header().Set("Content-Type", "image/jpeg")
	case "gif":
		gifImg, err := gif.DecodeAll(file)
		if err != nil {
			fmt.Println("Error decoding gif", err)
			return bytes.Buffer{}, err
		}
		err = gif.EncodeAll(&buf, gifImg)
		if err != nil {
			fmt.Println("Error encoding gif", err)
			return bytes.Buffer{}, err
		}
		w.Header().Set("Content-Type", "image/gif")
	default:
		fmt.Println("File type not supported")
		return bytes.Buffer{}, fmt.Errorf("File type not supported")
	}
	return buf, nil
}
