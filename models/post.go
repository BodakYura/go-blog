package models

type Post struct {
	Id string
	Title string
	Text string
}

func NewPost(id, title, text string) *Post  {
	return &Post{id, title, text}
}