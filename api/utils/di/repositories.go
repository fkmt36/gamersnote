package di

import (
	"gamersnote.com/v1/repositories/article"
	"gamersnote.com/v1/repositories/user"
)

var ArticleRepository article.Repository
var UserRepository user.Repository

func initRepo() {
	ArticleRepository = article.NewRepository(DBClient)
	UserRepository = user.NewRepository(DBClient)
}
