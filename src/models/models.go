package models

type Article struct {
	Id          int    `db:"article_id"`
	Title       string `db:"title"`
	Description string `db:"description"`
	Tags        []*Tag
}

type Tag struct {
	Id   int    `db:"tag_id"`
	Name string `db:"name"`
}
