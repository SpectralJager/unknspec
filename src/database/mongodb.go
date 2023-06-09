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
	client                *mongo.Client
	articlesCollection    *mongo.Collection
	adminTasksCollections *mongo.Collection
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
	tasks := client.Database("unknspec").Collection("adminTasks")

	return &MongoStorage{
		client:                client,
		articlesCollection:    articles,
		adminTasksCollections: tasks,
	}
}

func (db *MongoStorage) GetAllArticles(ctx context.Context) ([]*models.Article, error) {
	return db.filterArticles(ctx, bson.D{{}})
}

func (db *MongoStorage) GetArticlesByTitle(ctx context.Context, title string) ([]*models.Article, error) {
	return db.filterArticles(ctx, bson.D{{"title", bson.D{{"$regex", title}}}})
}

func (db *MongoStorage) GetOnlyPublishedArticles(ctx context.Context) ([]*models.Article, error) {
	return db.filterArticles(ctx, bson.D{{"is_draft", false}})
}

func (db *MongoStorage) GetArticle(ctx context.Context, id string) (*models.Article, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}
	articles, err := db.filterArticles(ctx, filter)
	if err != nil {
		return nil, err
	}
	if len(articles) == 0 {
		return nil, fmt.Errorf("article with id %s not found", id)
	}
	return articles[0], nil
}

func (db *MongoStorage) CreateArticle(ctx context.Context, article *models.Article) error {
	if article.Title == "" {
		return fmt.Errorf("article title cannot be empty")
	}
	if article.Abstract == "" {
		return fmt.Errorf("article abstract cannot be empty")
	}
	if article.Body == "" {
		return fmt.Errorf("article body cannot be empty")
	}
	filter := bson.D{{Key: "title", Value: article.Title}}
	articles, err := db.filterArticles(ctx, filter)
	if err != nil {
		return err
	}
	if articles != nil {
		return fmt.Errorf("article with title %s already defined", article.Title)
	}
	article.Id = primitive.NewObjectID()
	article.CreatedAt = time.Now().UTC()
	article.UpdatedAt = time.Now().UTC()
	_, err = db.articlesCollection.InsertOne(ctx, article)
	return err
}

func (db *MongoStorage) UpdateArticle(ctx context.Context, id string, article *models.Article) (*models.Article, error) {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}
	update := bson.D{{"$set", bson.D{
		{"title", article.Title},
		{"abstract", article.Abstract},
		{"body", article.Body},
		{"updated_at", time.Now().UTC()},
		{"is_draft", article.IsDraft},
		{"author", article.Author},
	}}}
	res, err := db.articlesCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}
	if res.ModifiedCount == 0 {
		return nil, fmt.Errorf("no articles was updated")
	}
	article.Id = objId
	return article, nil
}

func (db *MongoStorage) DeleteArticle(ctx context.Context, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}
	res, err := db.articlesCollection.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no articles were deleted")
	}
	return nil
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

func (db *MongoStorage) CreateAdminTask(ctx context.Context, task *models.AdminTask) error {
	task.Id = primitive.NewObjectID()
	_, err := db.adminTasksCollections.InsertOne(ctx, task)
	return err
}

func (db *MongoStorage) GetAdminTasks(ctx context.Context) ([]*models.AdminTask, error) {
	var tasks []*models.AdminTask
	cur, err := db.adminTasksCollections.Find(ctx, bson.D{{}})
	if err != nil {
		return tasks, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var task models.AdminTask
		err := cur.Decode(&task)
		if err != nil {
			return tasks, err
		}
		tasks = append(tasks, &task)
	}
	if err := cur.Err(); err != nil {
		return tasks, err
	}
	return tasks, nil
}

func (db *MongoStorage) DeleteAdminTask(ctx context.Context, id string) error {
	objId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{{Key: "_id", Value: objId}}
	res, err := db.adminTasksCollections.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	if res.DeletedCount == 0 {
		return fmt.Errorf("no task were deleted")
	}
	return nil
}
