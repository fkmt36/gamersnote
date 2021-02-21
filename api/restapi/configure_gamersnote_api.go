// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"gamersnote.com/v1/middlewares"
	"gamersnote.com/v1/restapi/operations"
	"gamersnote.com/v1/restapi/operations/healthcheck"
	"gamersnote.com/v1/utils/di"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
)

func configureFlags(api *operations.GamersnoteAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GamersnoteAPIAPI) http.Handler {
	api.ServeError = errors.ServeError

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.HealthcheckHealthcheckHandler = healthcheck.HealthcheckHandlerFunc(func(params healthcheck.HealthcheckParams) middleware.Responder {
		return healthcheck.NewHealthcheckOK()
	})

	// users
	api.UserPostUserHandler = di.PostUserHandler
	api.UserPatchUserVerifiedHandler = di.PatchVerifiedHandler
	api.UserPutUserHandler = di.PutMeHandler
	api.UserPatchUserEmailVerifiedHandler = di.PatchEmailVefifiedHandler
	api.UserGetUserHandler = di.GetUserHandler
	api.UserGetMeHandler = di.GetMeHandler
	api.UserPatchUserSigninedHandler = di.PatchSigninedHandler
	api.UserPatchUserSignoutedHandler = di.PatchSignoutedHandler
	api.UserPatchUserPasswordResetHandler = di.PatchPasswordResetHandler
	api.UserPutPasswordHandler = di.PutPasswordHandler
	api.UserDeleteUserHandler = di.DeleteMeHandler

	// image
	api.ImageUploadImageHandler = di.UploadImageHandler

	// article
	api.ArticlePostArticleHandler = di.PostArticleHandler
	api.ArticleGetArticlesHandler = di.GetAllArticlesHandler
	api.ArticleGetArticleHandler = di.GetArticleByArticleIDHandler
	api.ArticleGetTheUsersArticlesHandler = di.GetArticleByUsernameHandler
	api.ArticleDeleteArticleHandler = di.DeleteArticleHandler
	api.ArticlePutArticleHandler = di.PutArticleHandler
	api.ArticleGetLikedArticlesHandler = di.GetLikedArticleHandler

	// like
	api.LikeGetLikeHandler = di.GetLikeHandler
	api.LikePutLikeHandler = di.PutLikeHandler
	api.LikeDeleteLikeHandler = di.DeleteLikeHandler

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handler = middlewares.AccessLog(handler)
	handler = middlewares.Session(handler, di.SessionService, di.CtxuidService)
	return handler
}
