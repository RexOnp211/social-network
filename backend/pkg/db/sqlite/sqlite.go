package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"social-network/pkg/helpers"

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

func LoginUserDB(username string, password string) (helpers.Login, error) {
	var login helpers.Login
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

func AddPostToDb(data []interface{}) error {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return nil
	}

	stmt, err := DB.Prepare("INSERT INTO posts (post_id, user_id, subject, content, privacy, image) VALUES (?, ?, ?, ?, ?, ?)")
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

func GetPostsFromDb() ([]helpers.Post, error) {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return nil, err
	}

	rows, err := DB.Query("SELECT * FROM posts")
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	posts := []helpers.Post{}
	for rows.Next() {
		post := helpers.Post{}
		err := rows.Scan(&post.PostId, &post.UserId, &post.Subject, &post.Content, &post.CreationDate, &post.Image, &post.Privacy)
		if err != nil {
			log.Println("Scan error:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
