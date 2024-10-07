package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

// add file, file header and either "post" or "avatar" as imgtype
func SaveFile(file multipart.File, header *multipart.FileHeader, imgtype string) (string, error) {
	defer file.Close()
	var uploadDir string
	fmt.Println(header.Header)
	if imgtype == "post" {
		uploadDir = "./../../../frontend/public/image/upload"
	} else if imgtype == "avatar" {
		uploadDir = "./../../../frontend/public/image/avatar"
	}
	if _, err := os.Stat(uploadDir); os.IsNotExist(err) {
		os.Mkdir(uploadDir, os.ModePerm)
	}

	// TODO: add user id to filename
	userFileName := "userName" + "-" + header.Filename // userName is a placeholder
	filename := filepath.Join(uploadDir, userFileName)
	outFile, err := os.Create(filename)
	if err != nil {
		return "", err
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	var path string
	if imgtype == "post" {
		path = "/image/upload/" + userFileName
	} else if imgtype == "avatar" {
		path = "/image/avatar/" + userFileName
	}

	return path, nil
}
