package handlers

import (
	"fmt"
	"log"
	"net/http"
	"social-network/pkg"
	db "social-network/pkg/db/sqlite"

	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/register" {
		fmt.Println("404 not found.")
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		log.Println("Error at RegisterUser:", err)
		return
	}
	var User pkg.User

	log.Println("user:", User)
	// err := r.ParseMultipartForm(32 << 10)
	// if err != nil {
	// 	log.Println("Error at RegisterUser:", err)
	// 	return
	// }
	nickname := r.FormValue("nickname")
	email := r.FormValue("email")
	password := r.FormValue("password")
	firstname := r.FormValue("firstname")
	lastname := r.FormValue("lastname")
	dob := r.FormValue("dob")
	// avatar := r.FormFail"avatar");
	aboutMe := r.FormValue("aboutme")
	encryptedPassword, err2 := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err2 != nil {
		http.Error(w, "Error hashing password", http.StatusInternalServerError)
		return
	}

	log.Println("nickname:", nickname)
	log.Println("email:", email)
	log.Println("password:", password)
	log.Println("dob:", dob)

	data := []interface{}{nickname, email, string(encryptedPassword), firstname, lastname, dob, aboutMe}
	err = db.RegisterUserDB(data)
	log.Println("data:", data)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	fmt.Fprintf(w, "User created successfully!")
}
