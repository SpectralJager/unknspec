package db

import (
	"context"
	"testing"
	"time"
	"unknspec/src/models"
)

func TestGetDB(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db := NewSQLite("db.sqlite")
	if _, err := db.getDB(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGetArticle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db := NewSQLite("db.sqlite")
	if _, err := db.getDB(ctx); err != nil {
		t.Fatal(err)
	}
	article, err := db.ArticleGetById(0)
	if err != nil {
		t.Fatal(err)
	}
	if article.Id != 0 {
		t.Fatalf("Expected article with id %d, got %d", 0, article.Id)
	}
}

func TestGetArticleByTitle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db := NewSQLite("db.sqlite")
	if _, err := db.getDB(ctx); err != nil {
		t.Fatal(err)
	}
	article, err := db.ArticleGetByTitle("Test1")
	if err != nil {
		t.Fatal(err)
	}
	if article.Id != 0 {
		t.Fatalf("Expected article with id %d, got %d", 0, article.Id)
	}
}

func TestGetAllArticles(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db := NewSQLite("db.sqlite")
	if _, err := db.getDB(ctx); err != nil {
		t.Fatal(err)
	}
	_, err := db.ArticleAll()
	if err != nil {
		t.Fatal(err)
	}
}

func TestCreateArticle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db := NewSQLite("db.sqlite")
	if _, err := db.getDB(ctx); err != nil {
		t.Fatal(err)
	}
	_, err := db.ArticleCreate(&models.Article{Id: 0, Title: "foo", Description: "bar"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestUpdateArticle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db := NewSQLite("db.sqlite")
	if _, err := db.getDB(ctx); err != nil {
		t.Fatal(err)
	}
	article, err := db.ArticleGetByTitle("foo")
	if err != nil {
		t.Fatal(err)
	}
	_, err = db.ArticleUpdate(article.Id, &models.Article{Id: 0, Title: "fiz", Description: "baz"})
	if err != nil {
		t.Fatal(err)
	}
}

func TestDeleteArticle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	db := NewSQLite("db.sqlite")
	if _, err := db.getDB(ctx); err != nil {
		t.Fatal(err)
	}
	article, err := db.ArticleGetByTitle("fiz")
	if err != nil {
		t.Fatal(err)
	}
	err = db.ArticleDelete(article.Id)
	if err != nil {
		t.Fatal(err)
	}
}
