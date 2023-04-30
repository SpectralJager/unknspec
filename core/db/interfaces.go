package db

import "unknspec/core/models"

type PostApi interface {
	GetPosts() ([]models.Post, error)
	GetPostsByTitle(string) ([]models.Post, error)
	GetPostsByTags([]models.Tag) ([]models.Post, error)

	PostsPagination(int, int) ([]models.Post, error)

	GetPostById(int) (models.Post, error)
	CreatePost(map[string]any) (models.Post, error)
	UpdatePost(int, map[string]any) (models.Post, error)
	DeletePost(int) (models.Post, error)
}

type TagApi interface {
	GetTags() (models.Tag, error)
	GetTagsByName(string) (models.Tag, error)

	GetTagById(int) (models.Tag, error)
	CreateTag(map[string]any) (models.Tag, error)
	UpdateTag(map[string]any) (models.Tag, error)
	DeleteTag(int) (models.Tag, error)
}
