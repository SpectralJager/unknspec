package models

import (
	"time"
)

type Post struct {
	id       int    `json:"id"`
	title    string `json:"title"`
	abstract string `json:"abstract"`
	body     string `json:"body"`

	tags []Tag `json:"tags"`

	lastEditedAt time.Time `json:"last_edited_at"`
	isPublished  bool      `json:"is_published"`
}

func NewPost(title, abstract, body string, tags []Tag) (*Post, error) {
	var post Post

	post.title = title
	post.abstract = abstract
	post.body = body
	post.tags = tags
	post.lastEditedAt = time.Now()

	return &post, nil
}

func (p *Post) GetId() int {
	return p.id
}

func (p *Post) GetTitle() string {
	return p.title
}

func (p *Post) GetAbstract() string {
	return p.abstract
}

func (p *Post) GetBody() string {
	return p.body
}

func (p *Post) GetLastEditedAt() time.Time {
	return p.lastEditedAt
}

func (p *Post) IsPublished() bool {
	return p.isPublished
}

func (p *Post) TogglePublished() bool {
	p.isPublished = !p.isPublished
	return p.isPublished
}

func (p *Post) Update(title, abstract, body string, tags []Tag) error {
	p.title = title
	p.abstract = abstract
	p.body = body
	p.tags = tags
	p.lastEditedAt = time.Now()
	return nil
}
