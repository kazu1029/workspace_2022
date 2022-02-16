package model

type Article struct {
	ID int64 `json:"id"`
	Author string `json:"author"`
	Title string `json:"title"`
	Content string `json:"content"`
}