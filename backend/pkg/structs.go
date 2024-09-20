package pkg

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	PostBody string `json:"postBody"`
	Image    string `json:"image"`
	Privacy  string `json:"privacy"`
}
