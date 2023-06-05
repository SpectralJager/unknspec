package database

import (
	"context"
	"fmt"
	"unknspec/src/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	getDB(context.Context) (*sqlx.DB, error)
	ArticleGetById(context.Context, int) (*models.Article, error)
	ArticleGetByTitle(context.Context, string) (*models.Article, error)
	ArticleAll(context.Context, int, int) ([]*models.Article, error)

	TagGetById(context.Context, int) (*models.Tag, error)
	TagGetByName(context.Context, string) (*models.Tag, error)
	TagAll(context.Context) ([]*models.Tag, error)

	GetTagsOfArticleId(context.Context, int) ([]*models.Tag, error)
	GetArticlesFromTagsId(context.Context, ...int) ([]*models.Article, error)
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

func (db *SQLite) ArticleGetById(ctx context.Context, id int) (*models.Article, error) {
	articles := []*models.Article{}
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err := conn.SelectContext(ctx, &articles, "select * from articles where articles.article_id = ?", id); err != nil {
		return nil, err
	}
	return articles[0], nil
}

func (db *SQLite) ArticleGetByTitle(ctx context.Context, title string) (*models.Article, error) {
	articles := []*models.Article{}
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err := conn.SelectContext(ctx, &articles, "select * from articles where articles.title like $1", title); err != nil {
		return nil, err
	}
	return articles[0], nil
}

// if limit == 0, then return all articles
func (db *SQLite) ArticleAll(ctx context.Context, limit, offset int) ([]*models.Article, error) {
	articles := []*models.Article{}
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
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

func (db *SQLite) TagGetById(ctx context.Context, id int) (*models.Tag, error) {
	tags := []*models.Tag{}
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err := conn.SelectContext(ctx, &tags, "select * from tags where tag_id = ?", id); err != nil {
		return nil, err
	}
	return tags[0], nil
}

func (db *SQLite) TagGetByName(ctx context.Context, name string) (*models.Tag, error) {
	tags := []*models.Tag{}
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err := conn.SelectContext(ctx, &tags, "select * from tags where name like $1", name); err != nil {
		return nil, err
	}
	return tags[0], nil
}

// if limit == 0, then return all articles
func (db *SQLite) TagAll(ctx context.Context) ([]*models.Tag, error) {
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

func (db *SQLite) GetTagsOfArticleId(ctx context.Context, id int) ([]*models.Tag, error) {
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

func (db *SQLite) GetArticlesFromTagsId(ctx context.Context, ids ...int) ([]*models.Article, error) {
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	r := ""
	for _, id := range ids {
		r += fmt.Sprintf("%d,", id)
	}
	r = r[:len(r)-1]

	articles := []*models.Article{}
	if err := conn.SelectContext(ctx, &articles, fmt.Sprintf("select articles.* from articles inner join article_tag on article_tag.article_id = articles.article_id where article_tag.tag_id in (%s) group by articles.article_id", r)); err != nil {
		return nil, err
	}
	return articles, nil

}
