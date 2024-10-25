package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"social-network/pkg/helpers"
)

var DB *sql.DB

func SetDB(database *sql.DB) {
	DB = database
}

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

	// TODO: logic to use email when user input email

	err := DB.QueryRow("SELECT nickname, password FROM users WHERE nickname = ?", username).Scan(&login.Username, &login.Password)
	if err != nil {
		return login, errors.New("can't find username")
	}
	/* err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password))
	if err != nil {
		return login, errors.New("wrong Password")
	} */
	if login.Password != password {
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
		err := rows.Scan(&post.PostId, &post.UserId, &post.Subject, &post.Content, &post.Image, &post.Privacy, &post.CreationDate)
		if err != nil {
			log.Println("Scan error:", err)
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func GetUserFromDb(nickname string) (helpers.User, error) {
	user := helpers.User{}

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

func GetUserPostFromDbByUser(userId int) ([]helpers.Post, error) {
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
	posts := []helpers.Post{}
	for rows.Next() {
		post := helpers.Post{}
		err := rows.Scan(&post.PostId, &post.UserId, &post.Subject, &post.Content, &post.Image, &post.Privacy, &post.CreationDate)
		if err != nil {
			log.Println("Scan error:", err)
			return nil, err
		}
		posts = append(posts, post)
	}

	log.Println(posts)
	return posts, nil
}

func UpdateUserPrivacy(username string, privacyStatus string) {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
    if err != nil {
        fmt.Println("DB Open Error:", err)
        return
    }
    defer DB.Close()

	// string -> integer
	var publicStatus int
    if privacyStatus == "true" {
        publicStatus = 1
    } else if privacyStatus == "false" {
        publicStatus = 0
    } else {
        return
    }

    query := `UPDATE users SET public = ? WHERE nickname = ?`
    result, err := DB.Exec(query, publicStatus, username)
    if err != nil {
        log.Println("Update query error:", err)
        return
    }

	// show result of update
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        log.Println("Error checking affected rows:", err)
        return
    }
    if rowsAffected == 0 {
        log.Println("No user found with the given username:", username)
        return
    }

    log.Println("User privacy status updated successfully for username:", username)
    return
}

func GetGroupFromDb(groupname string) (helpers.Group, error) {
	group := helpers.Group{}

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
