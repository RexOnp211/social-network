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
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error:", err)
		return nil
	}

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

	rows, err := DB.Query("SELECT * FROM groups WHERE title = ?", groupname)
	if err != nil {
		log.Println("Query error:", err)
		return group, err
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&group.CreatorName, &group.Title, &group.Description)
		if err != nil {
			log.Println("Scan error:", err)
			return group, err
		}
	}

	return group, nil
}

func GetGroupsFromDb() ([]helpers.Group, error) {
	groups := []helpers.Group{}

	rows, err := DB.Query("SELECT * FROM groups")
	if err != nil {
		log.Println("Query error:", err)
		return groups, err
	}
	defer rows.Close()
	for rows.Next() {
		group := helpers.Group{}
		err := rows.Scan(&group.CreatorName, &group.Title, &group.Description)
		if err != nil {
			log.Println("Scan error:", err)
			return groups, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func CreateGroupDB(data []interface{}) error {

	stmt, err := DB.Prepare("INSERT INTO groups (creator_name, title, description) VALUES (?, ?, ?)")
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

type MembershipExistsError struct {
	Status string
}

func (e *MembershipExistsError) Error() string {
	return fmt.Sprintf("Membership already exists with status: %s", e.Status)
}

func InviteMemberDB(groupname string, username string, status string) (string, error) {

	var existingStatus string
	err := DB.QueryRow("SELECT status FROM group_members WHERE title = ? AND nickname = ?", groupname, username).Scan(&existingStatus)
	log.Println("TEST", existingStatus, err)

	// already there is data for the user & the group
	if existingStatus != "" {
		log.Println("Existing status found:", existingStatus)

		switch existingStatus {
		case "requested":
			log.Println("Case: requested")
			return fmt.Sprintf("User %s has already requested to join the group.", username), &MembershipExistsError{Status: existingStatus}
		case "invited":
			log.Println("Case: invited")
			return fmt.Sprintf("User %s has already been invited to the group.", username), &MembershipExistsError{Status: existingStatus}
		case "approved":
			log.Println("Case: approved")
			return fmt.Sprintf("User %s is already a member of the group.", username), &MembershipExistsError{Status: existingStatus}
		default:
			log.Println("Unexpected status encountered")
			return "Unexpected status.", nil
		}
	}

	// make new data
	stmt, err := DB.Prepare("INSERT INTO group_members (title, nickname, status) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Prepare statement error:", err)
		return "", err
	}
	defer stmt.Close()

	// the user does not exist
	_, err = stmt.Exec(groupname, username, status)
	if err != nil {
		if err.Error() == "FOREIGN KEY constraint failed" {
			return fmt.Sprintf("Invitation unsent: user %s does not exist", username), err
		}
		return "", err
	}

	return "", nil
}

func GetGroupMembersFromDb(nickname string) ([]helpers.GroupMembers, error) {
	rows, err := DB.Query("SELECT * FROM group_members WHERE nickname = ?", nickname)
	if err != nil {
		log.Println("Query error:", err)
		return nil, err
	}
	defer rows.Close()
	invitations := []helpers.GroupMembers{}
	for rows.Next() {
		invitation := helpers.GroupMembers{}
		err := rows.Scan(&invitation.Id, &invitation.Title, &invitation.Username, &invitation.Status)
		if err != nil {
			log.Println("Scan error:", err)
			return nil, err
		}
		invitations = append(invitations, invitation)
	}

	log.Println(invitations)
	return invitations, nil
}

func UpdateMemberStatus(id int, status string) error {
	var err error

	if status == "approve" {
		query := `UPDATE group_members SET status = ? WHERE id = ?`
		_, err = DB.Exec(query, "approved", id)
	}

	if status == "reject" {
		query := `DELETE FROM group_members WHERE id = ?`
		_, err = DB.Exec(query, id)
	}

	if err != nil {
		return fmt.Errorf("failed to update member status %v: %w", status, err)
	}
	return nil
}

func GetPostFromId(id int) (helpers.Post, error) {
	post := helpers.Post{}

	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in GetPostFromId:", err)
		return post, err
	}

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
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in GetCommentFromPostId:", err)
		return nil, err
	}

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
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in AddCommentToDb:", err)
		return err
	}

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
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in GetAvatarFromUserId:", err)
		return "", err
	}
	stmt := "SELECT avatar FROM users WHERE user_id = ?"
	avatar := ""
	err = DB.QueryRow(stmt, userId).Scan(&avatar)
	if err != nil {
		fmt.Println("QueryRow error in GetAvatarFromUserId:", err)
		return "", err
	}
	return avatar, nil
}

func GetNicknameFromId(id string) string {
	fmt.Println("id:", id)
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in GetNicknameFromId:", err)
		return ""
	}
	stmt := "SELECT nickname FROM users WHERE user_id = ?"
	nickname := ""
	err = DB.QueryRow(stmt, id).Scan(&nickname)
	if err != nil {
		fmt.Println("QueryRow error in GetNicknameFromId:", err)
		return ""
	}
	return nickname
}

func AddFollowRequestToDb(Fr helpers.FollowRequest) {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in AddFollowRequestToDb:", err)
		return
	}
	defer DB.Close()

	stmt, err := DB.Prepare("INSERT INTO followers (follower_id, followee_id, follows_back) VALUES (?, ?, ?)")
	if err != nil {
		log.Println("Prepare error in AddFollowRequestToDb:", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(Fr.FromUserId, Fr.ToUserId, Fr.FollowsBack)
	if err != nil {
		log.Println("Exec error in AddFollowRequestToDb:", err)
		return
	}
	return
}

func GetFollowRequestsFromDb(userId int) ([]helpers.FollowRequest, error) {
	DB, err := sql.Open("sqlite3", "../../pkg/db/database.db")
	if err != nil {
		fmt.Println("DB Open Error in GetFollowRequestsFromDb:", err)
		return nil, err
	}
	defer DB.Close()

	rows, err := DB.Query("SELECT follower_id, followee_id, follows_back FROM followers WHERE followee_id = ?", userId)
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
