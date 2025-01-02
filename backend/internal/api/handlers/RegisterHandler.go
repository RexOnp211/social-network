package handlers

import (
	"fmt"
	"log"
	"net/http"
	db "social-network/pkg/db/sqlite"
	"social-network/pkg/helpers"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		fmt.Println("404 not found.")
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		fmt.Println("Error parsing form", err)
		return
	}

	nickname := r.FormValue("nickname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	dob := r.FormValue("dob")
	file, header, err := r.FormFile("avatar")
	filepath := ""
	if err == nil {
		defer file.Close()
		filepath, err = helpers.SaveFile(file, header, "avatar")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			fmt.Println("Error saving file", err)
			return
		}
	}
	aboutMe := r.FormValue("aboutme")
	encryptedPassword, err2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err2 != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	data := []interface{}{nickname, email, string(encryptedPassword), firstname, lastname, dob, aboutMe, filepath}
	err = db.RegisterUserDB(data)
	log.Println("data:", data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprintf(w, "User created successfully!")
}
