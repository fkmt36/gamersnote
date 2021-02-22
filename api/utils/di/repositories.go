package di

import (
	"gamersnote.com/v1/repositories/article"
	"gamersnote.com/v1/repositories/comment"
	"gamersnote.com/v1/repositories/like"
	"gamersnote.com/v1/repositories/user"
)

var ArticleRepository article.Repository
var UserRepository user.Repository
var LikeRepository like.Repository
var CommentRepository comment.Repository

func initRepo() {
	ArticleRepository = article.NewRepository(DBClient)
	UserRepository = user.NewRepository(DBClient)
	LikeRepository = like.NewRepository(DBClient)
	CommentRepository = comment.NewRepository(DBClient)
}
