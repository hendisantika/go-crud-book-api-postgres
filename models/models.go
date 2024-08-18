package models

type Book struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishedAt string `json:"published_at"`
}
