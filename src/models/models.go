package models

type Article struct {
	Id          int    `db:"article_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
}
type Tag struct{}
