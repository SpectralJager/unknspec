package main

import "github.com/jmoiron/sqlx"

type Storage interface {
	getConn() (*sqlx.DB, error)

	GetPosts() ([]*Post, error)
	GetPostsRange(int, int) ([]*Post, error)
	GetPostsByTitle(string) ([]*Post, error)
	GetPostByID(int) (*Post, error)
	CreatePost(*Post) (*Post, error)
	UpdatePost(*Post, *Post) (*Post, error)
	DeletePost(*Post) (*Post, error)

	GetTags() ([]*Tag, error)
	GetTagsByName(string) ([]*Tag, error)
	GetTagByID(int) (*Tag, error)
	CreateTag(*Tag) (*Tag, error)
	UpdateTag(*Tag, *Tag) (*Tag, error)
	DeleteTag(*Tag) (*Tag, error)
}
