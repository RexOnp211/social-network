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

type ChatRoom struct {
	GroupId string `json:"groupId"`
}

type ChatRoomMembers struct {
	GroupId  string `json:"groupid"`
	Username string `json:"username"`
}

type PrivateMessage struct {
	GroupId    int    `json:"groupId"`
	FromUserId string `json:"fromUserId"`
	Content    string `json:"content"`
}

type GroupMembers struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Username string `json:"username"`
	Status   string `json:"status"`
}
type Membership struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	Username string `json:"username"`
	Status   string `json:"status"`
}

type GroupPost struct {
	Id           string    `json:"Id"`
	GroupTitle   string    `json:"groupTitle"`
	UserID       string    `json:"userId"`
	Nickname     string    `json:"nickname"`
	Subject      string    `json:"subject"`
	Content      string    `json:"content"`
	CreationDate string    `json:"creationDate"`
	Image        string    `json:"image"`
	Comments     []Comment `json:"comments"`
}

type GroupComment struct {
	CommentId    string `json:"commentId"`
	PostId       string `json:"postId"`
	UserID       string `json:"userId"`
	Nickname     string `json:"nickname"`
	Content      string `json:"content"`
	Image        string `json:"image"`
	CreationDate string `json:"creationDate"`
}

type GroupEvent struct {
	Id          int    `json:"Id"`
	GroupTitle  string `json:"groupTitle"`
	UserID      string `json:"userId"`
	Nickname    string `json:"nickname"`
	Title       string `json:"title"`
	Description string `json:"description"`
	EventDate   string `json:"eventDate"`
}

type EventStatus struct {
	Id       int    `json:"Id"`
	Nickname string `json:"nickname"`
	EventId  int    `json:"eventId"`
	Going    bool   `json:"going"`
}

type ChatMessage struct {
	From_id int    `json:"from"`
	To_id   int    `json:"to"`
	Message string `json:"message"`
}

type ChatRoom struct {
	ChatId  int   `json:"chatId`
	Members []int `json:"members"`
}
