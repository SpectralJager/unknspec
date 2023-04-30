package db

import "time"

type Post struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
	Body     string `json:"body"`

	CreatedAt   time.Time `json:"created_at"`
	PublishedAt time.Time `json:"published_at"`
	EditedAt    time.Time `json:"edited_at"`
}
