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

func (s *PostgresStorage) InitTables() error {
	schemas := []string{
		`create table if not exists posts (
			id serial primary key,
			title varchar(128) not null unique,
			abstract varchar(512) not null,
			body text not null,
			edited_at timestamp not null,
			published bool default false
		)`,
		`create table if not exists tags (
			id serial primary key,
			name varchar(64) not null unique
		)`,
		`create table if not exists posts_tags (
			post_id int not null,
			tag_id int not null,
			primary key(post_id, tag_id),
			foreign key(post_id) 
				references posts(id)
				on delete cascade,		
			foreign key(tag_id) 
				references tags(id)
				on delete cascade
		)`,
	}
	for _, sch := range schemas {
		db, err := s.getConn()
		if err != nil {
			return err
		}
		_, err = db.Exec(sch)
		if err != nil {
			return err
		}
	}
	return nil
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
