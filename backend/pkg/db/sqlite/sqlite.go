package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"social-network/pkg"

	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func RegisterUserDB(data []interface{}) error {

	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return nil
	}

	stmt, err := DB.Prepare("INSERT INTO users (nickname, email, password, firstname, lastname, dob, aboutme, avatar) VALUES (?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data...)
	if err != nil {
		log.Println("Exec statement error:", err)
		return err
	}
	return nil
}

func LoginUserDB(username string, password string) (pkg.Login, error) {
	var login pkg.Login
	var fieldname string

	err := DB.QueryRow("SELECT username, password FROM users WHERE "+fieldname+" = ?", username).Scan(&login.Username, &login.Password)
	if err != nil {
		return login, errors.New("can't find username")
	}
	err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password))
	if err != nil {
		return login, errors.New("wrong Password")
	}
	return login, nil
}

// GetUserIDByUsernameOrEmail retrieves a user ID based on their username or email.
func GetUserIDByUsernameOrEmail(username string) (int, error) {
	var userID int
	var fieldname string

	err := DB.QueryRow("SELECT user_id FROM users WHERE "+fieldname+" = ?", username).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
