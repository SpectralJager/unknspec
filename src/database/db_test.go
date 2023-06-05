package database

import (
	"context"
	"testing"
	"time"
)

var db = NewSQLite("db.sqlite")

func TestGetDB(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	if _, err := db.getDB(ctx); err != nil {
		t.Fatal(err)
	}
}

func TestGetArticle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	article, err := db.ArticleGetById(ctx, 0)
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
	article, err := db.ArticleGetByTitle(ctx, "Test1")
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
	_, err := db.ArticleAll(ctx, 0, 0)
	if err != nil {
		t.Fatal(err)
	}
	articles, err := db.ArticleAll(ctx, 2, 1)
	if err != nil {
		t.Fatal(err)
	}
	if len(articles) != 2 {
		t.Fatalf("expected 2 records, get %d", len(articles))
	}
}

func TestGetTag(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tag, err := db.TagGetById(ctx, 0)
	if err != nil {
		t.Fatal(err)
	}
	if tag.Id != 0 {
		t.Fatalf("Expected tag with id %d, got %d", 0, tag.Id)
	}
}

func TestGetTagByName(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tag, err := db.TagGetByName(ctx, "tag1")
	if err != nil {
		t.Fatal(err)
	}
	if tag.Id != 0 {
		t.Fatalf("Expected tag with id %d, got %d", 0, tag.Id)
	}
}

func TestGetAllTags(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := db.TagAll(ctx)
	if err != nil {
		t.Fatal(err)
	}
}

func TestGetTagsOfArticle(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	tags, err := db.GetTagsOfArticleId(ctx, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, val := range tags {
		t.Logf("got %v", val)
	}
	// t.Fatal()
}

func TestGetArticlesByTags(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()
	articles, err := db.GetArticlesFromTagsId(ctx, 0)
	if err != nil {
		t.Fatal(err)
	}
	for _, val := range articles {
		t.Logf("got %v", val)
	}
	// t.Fatal()
}
