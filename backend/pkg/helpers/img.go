package helpers

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
)

// add file, file header and either "post" or "avatar" as imgtype
func SaveFile(file multipart.File, header *multipart.FileHeader, imgtype string) (string, error) {
	defer file.Close()
	uploadDir := "../../assets/image/"
	fmt.Println(header.Header)
	if imgtype == "post" {
		uploadDir += "upload/"
	} else if imgtype == "avatar" {
		uploadDir += "avatar/"
	}
	_, err := os.Stat(uploadDir)
	if os.IsNotExist(err) {
		err := os.MkdirAll(uploadDir, 0755)
		if err != nil {
			return "", err
		}
	}

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

	_, err = io.Copy(outFile, file)
	if err != nil {
		return "", err
	}

	return fileName, nil
}
