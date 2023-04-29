package db

type PostAccess interface {
	GetPosts() []Post
	GetPostById(int) Post
	GetPostByTitle() []Post
}
