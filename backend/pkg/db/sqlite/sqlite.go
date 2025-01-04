package db

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"

	"social-network/pkg/helpers"

	"golang.org/x/crypto/bcrypt"
)

var DB *sql.DB

func SetDB(database *sql.DB) {
	DB = database
}

func RegisterUserDB(data []interface{}) error {

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

	err := DB.QueryRow("SELECT user_id, nickname, password FROM users WHERE nickname = ?", username).Scan(&login.UserId, &login.Username, &login.Password)
	if err != nil {
		return login, errors.New("can't find username")
	}
	err = bcrypt.CompareHashAndPassword([]byte(login.Password), []byte(password))
	if err != nil {
		return login, errors.New("wrong Password")
	}
	// fmt.Println("login.Password", login.Password, "password", password)
	// if login.Password != password {
	// 	return login, errors.New("wrong Password")
	// }
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

	stmt, err := DB.Prepare("INSERT INTO posts (user_id, subject, content, image, privacy) VALUES (?, ?, ?, ?, ?)")
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

func GetUserPostFromDbByUser(userId string) ([]helpers.Post, error) {
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
}

func GetPostFromId(id int) (helpers.Post, error) {
	post := helpers.Post{}

	rows, err := DB.Query("SELECT * FROM posts WHERE post_id = ?", id)
	if err != nil {
		log.Println("Query error in GetPostFromId:", err)
		return post, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&post.PostId, &post.UserId, &post.Subject, &post.Content, &post.Image, &post.Privacy, &post.CreationDate)
		if err != nil {
			log.Println("Scan error in GetPostFromId:", err)
			return post, err
		}
	}

	return post, nil
}

func GetCommentsFromPostId(id int) ([]helpers.Comment, error) {

	rows, err := DB.Query("SELECT comment_id, post_id, user_id, content, image FROM comments WHERE post_id = ?", id)
	if err != nil {
		log.Println("Query error in GetCommentFromPostId:", err)
		return nil, err
	}
	defer rows.Close()
	comments := []helpers.Comment{}
	for rows.Next() {
		comment := helpers.Comment{}
		err := rows.Scan(&comment.CommentId, &comment.PostId, &comment.UserId, &comment.Content, &comment.Image)
		if err != nil {
			log.Println("Scan error in GetCommentFromPostId:", err)
			return nil, err
		}
		comments = append(comments, comment)
	}

	return comments, nil
}

func AddCommentToDb(data []interface{}) error {
	fmt.Println("interfacedata", data)

	stmt, err := DB.Prepare("INSERT INTO comments (post_id, user_id, content, image) VALUES (?, ?, ?, ?)")
	if err != nil {
		log.Println("Prepare error in AddCommentToDb:", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(data...)
	if err != nil {
		log.Println("Exec error in AddCommentToDb:", err)
		return err
	}
	return nil
}

func GetAvatarFromUserId(userId string) (string, error) {

	stmt := "SELECT avatar FROM users WHERE user_id = ?"
	avatar := ""
	err := DB.QueryRow(stmt, userId).Scan(&avatar)
	if err != nil {
		fmt.Println("QueryRow error in GetAvatarFromUserId:", err)
		return "", err
	}
	return avatar, nil
}

func GetNicknameFromId(id string) string {
	fmt.Println("id:", id)

	stmt := "SELECT nickname FROM users WHERE user_id = ?"
	nickname := ""
	err := DB.QueryRow(stmt, id).Scan(&nickname)
	if err != nil {
		fmt.Println("QueryRow error in GetNicknameFromId:", err)
		return ""
	}
	return nickname
}

func AddFollowRequestToDb(Fr helpers.FollowRequest) error {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in AddFollowRequestToDb:", err)
		return err
	}
	defer DB.Close()

	stmt2, err := DB.Prepare("INSERT INTO followers (follower_id, followee_id, accepted) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Prepare error in AddFollowRequestToDb:", err)
		return err
	}
	defer stmt2.Close()

	_, err = stmt2.Exec(Fr.FromUserId, Fr.ToUserId, Fr.FollowsBack)
	if err != nil {
		log.Println("Exec error in AddFollowRequestToDb:", err)
		return err
	}
	return nil
}

func GetFollowRequestsFromDb(userId int) ([]helpers.FollowRequest, error) {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in GetFollowRequestsFromDb:", err)
		return nil, err
	}
	defer DB.Close()

	rows, err := DB.Query("SELECT follower_id, followee_id, accepted FROM followers WHERE followee_id = ? AND accepted = ?", userId, false)
	if err != nil {
		log.Println("Query error in GetFollowRequestsFromDb:", err)
		return nil, err
	}
	defer rows.Close()
	followRequests := []helpers.FollowRequest{}
	for rows.Next() {
		fr := helpers.FollowRequest{}
		err := rows.Scan(&fr.FromUserId, &fr.ToUserId, &fr.FollowsBack)
		if err != nil {
			log.Println("Scan error in GetFollowRequestsFromDb:", err)
			return nil, err
		}
		followRequests = append(followRequests, fr)
	}
	return followRequests, nil
}

func UpdateFollowRequestStatusDB(from, to string, status bool) error {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in GetFollowRequestsFromDb:", err)
		return err
	}
	defer DB.Close()

	var querry string

	// changes querry based on if followrquest accepted or declined
	if status {
		querry = "UPDATE followers SET accepted = true WHERE follower_id = ? AND followee_id = ?"
	} else {
		querry = "DELETE FROM followers WHERE follower_id = ? AND followee_id ?"
		stmt, err := DB.Prepare(querry)
		if err != nil {
			log.Println("Error Changing follow status", err)
			return err
		}
		defer stmt.Close()

		_, err2 := stmt.Exec(from, to)
		if err2 != nil {
			log.Println("Error executing followRequest Accept in db", err)
			return err
		}
	}

	return nil
}

func GetUsersFollowingListFromDb(userId int) ([]helpers.User, error) {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in GetFollowRequestsFromDb:", err)
		return nil, err
	}
	defer DB.Close()

	rows, err := DB.Query("SELECT followee_id FROM followers WHERE follower_id = ?", userId)
	if err != nil {
		fmt.Println("error querrying db for follwee ids", err)
		return nil, err
	}

	followingUsersArr := []int{}
	for rows.Next() {
		var followingId int
		err := rows.Scan(&followingId)
		if err != nil {
			fmt.Println("error scanning db for followee id")
			return nil, err
		}
		followingUsersArr = append(followingUsersArr, followingId)
	}

	userArr := []helpers.User{}
	for _, id := range followingUsersArr {
		strId := strconv.Itoa(id)
		fmt.Println("THIS IS FOLLOWING ID", id)

		nickname := GetNicknameFromId(strId)
		user, err := GetUserFromDb(nickname)
		if err != nil {
			fmt.Println("Error getting user for sidebar")
			return nil, err
		}
		userArr = append(userArr, user)
	}

	return userArr, nil
}

func GetUsersFollowersListFromDB(userId int) ([]helpers.User, error) {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in GetFollowRequestsFromDb:", err)
		return nil, err
	}
	defer DB.Close()

	rows, err := DB.Query("SELECT follower_id FROM followers WHERE followee_id = ?", userId)
	if err != nil {
		fmt.Println("error querrying db for follwer ids", err)
		return nil, err
	}

	followersUsersArr := []int{}
	for rows.Next() {
		var followerId int
		err := rows.Scan(&followerId)
		if err != nil {
			fmt.Println("error scanning db for follower id")
			return nil, err
		}
		followersUsersArr = append(followersUsersArr, followerId)
	}

	userArr := []helpers.User{}
	for _, id := range followersUsersArr {
		fmt.Println("THIS IS ID FOLLOWERS", id)
		strId := strconv.Itoa(id)

		nickname := GetNicknameFromId(strId)
		user, err := GetUserFromDb(nickname)
		if err != nil {
			fmt.Println("error getting user for profilepage")
			return nil, err
		}
		userArr = append(userArr, user)
	}

	return userArr, nil
}
