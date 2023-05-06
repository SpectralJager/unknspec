package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresStorage struct {
	ConnString string
}

func NewPostgresStorage(connStr string) *PostgresStorage {
	return &PostgresStorage{
		ConnString: connStr,
	}
}

func (s *PostgresStorage) getConn() (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", s.ConnString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (s *PostgresStorage) GetPosts() ([]*Post, error) {
	return nil, nil
}

func (s *PostgresStorage) GetPostsRange(from, to int) ([]*Post, error) {
	return nil, nil
}

func (s *PostgresStorage) GetPostsByTitle(title string) ([]*Post, error) {
	return nil, nil
}

func (s *PostgresStorage) GetPostByID(id int) (*Post, error) {
	return nil, nil
}

func (s *PostgresStorage) CreatePost(post *Post) (*Post, error) {
	return nil, nil
}

func (s *PostgresStorage) UpdatePost(old, new *Post) (*Post, error) {
	return nil, nil
}

func (s *PostgresStorage) DeletePost(post *Post) (*Post, error) {
	return nil, nil
}

func (s *PostgresStorage) GetTags() ([]*Tag, error) {
	return nil, nil
}

func (s *PostgresStorage) GetTagsByName(name string) ([]*Tag, error) {
	return nil, nil
}

func (s *PostgresStorage) GetTagByID(id int) (*Tag, error) {
	return nil, nil
}

func (s *PostgresStorage) CreateTag(tag *Tag) (*Tag, error) {
	return nil, nil
}

func (s *PostgresStorage) UpdateTag(old, new *Tag) (*Tag, error) {
	return nil, nil
}

func (s *PostgresStorage) DeleteTag(tag *Tag) (*Tag, error) {
	return nil, nil
}
