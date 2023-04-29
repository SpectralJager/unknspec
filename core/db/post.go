package db

import "time"

type Post struct {
	Id       int
	Title    string
	Abstract string
	Body     string

	CreatedAt   time.Time
	PublishedAt time.Time
	EditedAt    time.Time
}
