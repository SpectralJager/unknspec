package database

import (
	"context"
	"fmt"
	"log"
	"time"
	"unknspec/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoStorage struct {
	client             *mongo.Client
	articlesCollection *mongo.Collection
}

func NewMongoStorage(url string) *MongoStorage {
	ctx := context.TODO()
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		log.Fatal(err)
	}
	articles := client.Database("unknspec").Collection("articles")

	return &MongoStorage{
		client:             client,
		articlesCollection: articles,
	}
}

func (db *MongoStorage) GetAllArticles(ctx context.Context) ([]*models.Article, error) {
	return db.filterArticles(ctx, bson.D{{}})
}

func (db *MongoStorage) GetArticlesWithTag(ctx context.Context, tag string) ([]*models.Article, error) {
	filter := bson.D{{"tags", bson.D{{"$all", bson.A{tag}}}}}
	return db.filterArticles(ctx, filter)
}

func (db *MongoStorage) CreateArticle(ctx context.Context, article *models.Article) error {
	filter := bson.D{
		{Key: "title", Value: article.Title},
	}
	articles, err := db.filterArticles(ctx, filter)
	if err != nil {
		return err
	}
	if articles != nil {
		return fmt.Errorf("article with title %s already defined", article.Title)
	}
	article.Id = primitive.NewObjectID()
	article.CreatedAt = time.Now().UTC()
	article.UpadtedAt = time.Now().UTC()
	_, err = db.articlesCollection.InsertOne(ctx, article)
	return err
}

func (db *MongoStorage) filterArticles(ctx context.Context, filter bson.D) ([]*models.Article, error) {
	var articles []*models.Article
	cur, err := db.articlesCollection.Find(ctx, filter)
	if err != nil {
		return articles, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var article models.Article
		err := cur.Decode(&article)
		if err != nil {
			return articles, err
		}
		articles = append(articles, &article)
	}
	if err := cur.Err(); err != nil {
		return articles, err
	}
	return articles, nil
}
