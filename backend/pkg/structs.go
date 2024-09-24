package pkg

type Post struct {
	PostId       string `json:"postId"`
	UserId       string `json:"userId"`
	Subject      string `json:"subject"`
	Content      string `json:"content"`
	Privacy      string `json:"privacy"`
	CreationDate string `json:"creationDate"`
	Image        string `json:"image"`
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
}
