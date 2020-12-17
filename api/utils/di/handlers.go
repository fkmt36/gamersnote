package di

import (
	"gamersnote.com/v1/handlers/article"
	"gamersnote.com/v1/handlers/image"
	"gamersnote.com/v1/handlers/user"
)

// article
var GetAllArticlesHandler *article.GetAllArticlesHandler
var GetArticleByArticleIDHandler *article.GetArticleByArticleIDHandler
var PostArticleHandler *article.PostArticleHandler

// image
var UploadImageHandler *image.UploadImageHandler

// user
var DeleteMeHandler *user.DeleteMeHandler
var GetMeHandler *user.GetMeHandler
var GetUserHandler *user.GetUserHandler
var PatchSigninedHandler *user.PatchSigninedHandler
var PatchVerifiedHandler *user.PatchVefifiedHandler
var PostUserHandler *user.PostUserHandler
var PutMeHandler *user.PutMeHandler
var PatchSignoutedHandler *user.PatchSignoutedHandler

func initHdlr() {
	GetAllArticlesHandler = article.NewGetAllArticlesHandler(ArticleRepository)
	GetArticleByArticleIDHandler = article.NewGetArticleByArticleIDHandler(ArticleRepository)
	PostArticleHandler = article.NewPostArticleHandler(ArticleRepository, CtxuidService)
	UploadImageHandler = image.NewUploadImageHandler(ImageService)
	DeleteMeHandler = user.NewDeleteMeHandler(UserRepository)
	GetMeHandler = user.NewGetMeHandler(UserRepository, CtxuidService)
	GetUserHandler = user.NewGetUserHandler(UserRepository)
	PatchSigninedHandler = user.NewPatchSigninedHandler(UserRepository, SessionService)
	PatchVerifiedHandler = user.NewPatchVefifiedHandler(UserRepository, TmpuserService, SessionService)
	PostUserHandler = user.NewPostUserHandler(UserRepository, TmpuserService, EmailSender)
	PutMeHandler = user.NewPutMeHandler(UserRepository)
	PatchSignoutedHandler = user.NewPatchSignoutedHandler(CtxuidService, SessionService)
}
