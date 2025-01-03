package helpers

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

// add file, file header and either "post" or "avatar" as imgtype
func SaveFile(file multipart.File, header *multipart.FileHeader, imgtype string) (string, error) {
	log.Println("posting to group...")

	defer file.Close()
	uploadDir := "../../assets/image/"
	fmt.Println("header", header.Header)
	if imgtype == "post" {
		uploadDir += "upload/"
	} else if imgtype == "avatar" {
		uploadDir += "avatar/"
	} else if imgtype == "group-post" {
		uploadDir += "group-post/"
	} else if imgtype == "group-post-comment" {
		uploadDir += "group-post-comment/"
	}
	_, err := os.Stat(uploadDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, 0755)
		if err != nil {
			return "", err
		}
	}
	log.Println("posting to group...")

	// TODO: add image id to filename
	files, err := os.ReadDir(uploadDir)
	if err != nil {
		return "", err
	}
	imgId := len(files) + 1
	fileName := strconv.Itoa(imgId) + "_" + header.Filename
	filepath := filepath.Join(uploadDir, fileName)
	outFile, err := os.Create(filepath)
	if err != nil {
		return "", err
	}
	defer outFile.Close()
	log.Println("posting to group...")


	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
