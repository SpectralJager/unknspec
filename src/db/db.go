package db

import (
	"context"
	"time"
	"unknspec/src/models"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Database interface {
	getDB(context.Context) (*sqlx.DB, error)
	ArticleGetById(int) (*models.Article, error)
	ArticleGetByTitle(string) (*models.Article, error)
	ArticleAll() ([]*models.Article, error)
	ArticleDelete(int) error
	ArticleUpdate(int, *models.Article) (*models.Article, error)
	ArticleCreate(*models.Article) (*models.Article, error)
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

func (db *SQLite) ArticleGetById(id int) (*models.Article, error) {
	articles := []*models.Article{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
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

func (db *SQLite) ArticleGetByTitle(title string) (*models.Article, error) {
	articles := []*models.Article{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
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

func (db *SQLite) ArticleAll() ([]*models.Article, error) {
	articles := []*models.Article{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if err := conn.SelectContext(ctx, &articles, "select * from articles"); err != nil {
		return nil, err
	}
	return articles, nil
}

func (db *SQLite) ArticleDelete(id int) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := db.getDB(ctx)
	if err != nil {
		return err
	}
	defer conn.Close()
	if _, err := conn.ExecContext(ctx, "delete from articles where article_id = ?", id); err != nil {
		return err
	}
	return nil
}

func (db *SQLite) ArticleUpdate(id int, article *models.Article) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	article.Id = id
	if _, err := conn.NamedExecContext(ctx, "update articles set title=:title, description=:description where article_id = :article_id", article); err != nil {
		return nil, err
	}
	return article, nil
}
func (db *SQLite) ArticleCreate(article *models.Article) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	conn, err := db.getDB(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	if _, err := conn.NamedExecContext(ctx, "insert into articles (title, description) values (:title, :description)", article); err != nil {
		return nil, err
	}
	return article, nil
}
