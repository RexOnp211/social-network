package helpers

import (
	"time"
)

type Post struct {
	PostId       string    `json:"postId"`
	UserId       string    `json:"userId"`
	Subject      string    `json:"subject"`
	Content      string    `json:"content"`
	Privacy      string    `json:"privacy"`
	CreationDate string    `json:"creationDate"`
	Image        string    `json:"image"`
	Comments     []Comment `json:"comments"`
}

type User struct {
	Id        int    `json:"id"`
	Nickname  string `json:"nickname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Dob       string `json:"dob"`
	AboutMe   string `json:"aboutMe"`
	Public    bool   `json:"public"`
	Avatar    string `json:"avatar"`
}

type Login struct {
	UserId   int    `json:"id"`
	Username string `json:"email"`
	Password string `json:"password"`
}

type Session struct {
	Username     string    `json:"username"`
	SessionToken string    `json:"session_token"`
	ExpireTime   time.Time `json:"expire_time"`
	UserID       int       `json:"user_id"`
}

type Comment struct {
	CommentId    string `json:"commentId"`
	PostId       string `json:"postId"`
	UserId       string `json:"userId"`
	Content      string `json:"content"`
	Image        string `json:"image"`
	CreationDate string `json:"creationDate"`
}
type Group struct {
	CreatorName string `json:"creatorName"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type FollowRequest struct {
	FromUserId   string `json:"fromUserId"`
	ToUserId     string `json:"toUserId"`
	FollowsBack  bool   `json:"followsBack"`
	CreationDate string `json:"creationDate"`
}

type GroupMembers struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Username string `json:"username"`
	Status   string `json:"status"`
}
