package database

import (
	"context"
	"unknspec/src/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	getDB(context.Context) (*sqlx.DB, error)

	GetArticleById(context.Context, int) (*models.Article, error)
	GetArticleByTitle(context.Context, string) (*models.Article, error)
	GetAllArticles(context.Context, int, int) ([]*models.Article, error)

	GetTagById(context.Context, int) (*models.Tag, error)
	GetTagByName(context.Context, string) (*models.Tag, error)
	GetAllTags(context.Context) ([]*models.Tag, error)

	GetTagsForArticle(context.Context, int) ([]*models.Tag, error)
	GetArticlesForTag(context.Context, int) ([]*models.Article, error)

	CreateArticle(context.Context, *models.Article) error
	CreateTags(context.Context, *models.Tag) error

	AddTagToArticle(context.Context, int, int) error

	UpdateArticle(context.Context, *models.Article) error
	UpdateTag(context.Context, *models.Tag) error

	DeleteArticle(context.Context, int) error
	DeleteTag(context.Context, int) error
	RemoveTagFromArticle(context.Context, int, int) error
}

type SQLite struct {
	filename string
}

func NewSQLite(filename string) Database {
	return &SQLite{
		filename: filename,
	}
}

func (db *SQLite) getDB(ctx context.Context) (*sqlx.DB, error) {
	conn, err := sqlx.ConnectContext(ctx, "sqlite3", db.filename)
	if err != nil {
		return conn, err
	}
	if err = conn.Ping(); err != nil {
		return nil, err
	}
	return conn, nil
}

func (db *SQLite) GetArticleById(ctx context.Context, id int) (*models.Article, error) {
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	articles := []*models.Article{}
	if err := conn.SelectContext(ctx, &articles, "select * from articles where articles.article_id = $1", id); err != nil {
		return nil, err
	}
	return articles[0], nil
}

func (db *SQLite) GetArticleByTitle(ctx context.Context, title string) (*models.Article, error) {
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	articles := []*models.Article{}
	if err := conn.SelectContext(ctx, &articles, "select * from articles where articles.title like $1", title); err != nil {
		return nil, err
	}
	return articles[0], nil
}

// if limit == 0, then return all articles
func (db *SQLite) GetAllArticles(ctx context.Context, limit, offset int) ([]*models.Article, error) {
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	articles := []*models.Article{}
	if limit == 0 {
		if err := conn.SelectContext(ctx, &articles, "select * from articles"); err != nil {
			return nil, err
		}
	} else {
		if err := conn.SelectContext(ctx, &articles, "select * from articles limit $1 offset $2", limit, offset); err != nil {
			return nil, err
		}
	}
	return articles, nil
}

func (db *SQLite) GetTagById(ctx context.Context, id int) (*models.Tag, error) {
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tags := []*models.Tag{}
	if err := conn.SelectContext(ctx, &tags, "select * from tags where tag_id = ?", id); err != nil {
		return nil, err
	}
	return tags[0], nil
}

func (db *SQLite) GetTagByName(ctx context.Context, name string) (*models.Tag, error) {
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tags := []*models.Tag{}
	if err := conn.SelectContext(ctx, &tags, "select * from tags where name like $1", name); err != nil {
		return nil, err
	}
	return tags[0], nil
}

// if limit == 0, then return all articles
func (db *SQLite) GetAllTags(ctx context.Context) ([]*models.Tag, error) {
	tags := []*models.Tag{}
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err := conn.SelectContext(ctx, &tags, "select * from tags"); err != nil {
		return nil, err
	}
	return tags, nil
}

func (db *SQLite) GetTagsForArticle(ctx context.Context, id int) ([]*models.Tag, error) {
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	tags := []*models.Tag{}
	if err := conn.SelectContext(ctx, &tags, "SELECT tags.tag_id, tags.name FROM tags inner join article_tag on article_tag.tag_id = tags.tag_id where article_tag.article_id = $1;", id); err != nil {
		return nil, err
	}
	return tags, nil
}

func (db *SQLite) GetArticlesForTag(ctx context.Context, id int) ([]*models.Article, error) {
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	articles := []*models.Article{}
	if err := conn.SelectContext(ctx, &articles, "select articles.* from articles inner join article_tag on article_tag.article_id = articles.article_id where article_tag.tag_id = $1 group by articles.article_id", id); err != nil {
		return nil, err
	}
	return articles, nil

}

func (db *SQLite) CreateArticle(ctx context.Context, article *models.Article) error {
	conn, err := db.getDB(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err := conn.NamedExecContext(ctx, "insert into articles (title, description) values (:title,:description)", article); err != nil {
		return err
	}
	return nil
}

func (db *SQLite) CreateTags(ctx context.Context, tag *models.Tag) error {
	conn, err := db.getDB(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err := conn.NamedExecContext(ctx, "insert into tags (name) values (:name)", tag); err != nil {
		return err
	}
	return nil
}

func (db *SQLite) AddTagToArticle(context.Context, int, int) error {
	return nil
}

func (db *SQLite) UpdateArticle(context.Context, *models.Article) error {
	return nil
}

func (db *SQLite) UpdateTag(context.Context, *models.Tag) error {
	return nil
}

func (db *SQLite) DeleteArticle(context.Context, int) error {
	return nil
}

func (db *SQLite) DeleteTag(context.Context, int) error {
	return nil
}

func (db *SQLite) RemoveTagFromArticle(context.Context, int, int) error {
	return nil
}
