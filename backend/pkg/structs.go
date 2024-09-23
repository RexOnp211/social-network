package pkg

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	PostBody string `json:"postBody"`
	Image    string `json:"image"`
	Privacy  string `json:"privacy"`
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
