package main

import "time"

type Post struct {
	ID        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Abstract  string    `json:"abstract" db:"abstract"`
	Tags      []Tag     `json:"tags" db:"tags"`
	Body      string    `json:"body" db:"body"`
	EditedAt  time.Time `json:"edited_at" db:"edited_at"`
	Published bool      `json:"published" db:"published"`
}

func NewPost(title, abstract, body string, tags []Tag, published bool) *Post {
	return &Post{
		Title:     title,
		Abstract:  abstract,
		Tags:      tags,
		Body:      body,
		Published: published,
		EditedAt:  time.Now().UTC(),
	}
}

type Tag struct {
	ID   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}

func NewTag(name string) *Tag {
	return &Tag{
		Name: name,
	}
}
