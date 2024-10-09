package db

import (
	"database/sql"
	"fmt"
	"log"
	"social-network/pkg"
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

func GetPostsFromDb() ([]pkg.Post, error) {
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
	posts := []pkg.Post{}
	for rows.Next() {
		post := pkg.Post{}
		err := rows.Scan(&post.PostId, &post.UserId, &post.Subject, &post.Content, &post.CreationDate, &post.Image, &post.Privacy)
		if err != nil {
			log.Println("Scan error:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetUserFromDb(nickname string) (pkg.User, error) {
	user := pkg.User{}

	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return user, err
	}

	rows, err := DB.Query("SELECT user_id, nickname, email, firstname, lastname, dob, aboutme, public, avatar FROM users WHERE nickname = ?", nickname)
	if err != nil {
		log.Println("Query error:", err)
		return user, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Nickname, &user.Email, &user.Firstname, &user.Lastname, &user.Dob, &user.AboutMe, &user.Public, &user.Avatar)
		if err != nil {
			log.Println("Scan error:", err)
			return user, err
		}
	}

	return user, nil
}

func GetUserPostFromDbByUser(userId int) ([]pkg.Post, error) {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return nil, err
	}

	rows, err := DB.Query("SELECT * FROM posts WHERE user_id = ?", userId)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	posts := []pkg.Post{}
	for rows.Next() {
		post := pkg.Post{}
		err := rows.Scan(&post.PostId, &post.UserId, &post.Subject, &post.Content, &post.CreationDate, &post.Image, &post.Privacy)
		if err != nil {
			log.Println("Scan error:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	log.Println(posts)
	return posts, nil
}

func GetGroupFromDb(groupname string) (pkg.Group, error) {
	group := pkg.Group{}

	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return group, err
	}

	rows, err := DB.Query("SELECT * FROM groups WHERE title = ?", groupname)
	if err != nil {
		log.Println("Query error:", err)
		return group, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&group.Id, &group.CreatorId, &group.Title, &group.Description)
		if err != nil {
			log.Println("Scan error:", err)
			return group, err
		}
	}

	return group, nil
}