// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"gamersnote.com/v1/restapi/operations/article"
	"gamersnote.com/v1/restapi/operations/user"

	"gamersnote.com/v1/utils"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"gamersnote.com/v1/configs"
	"gamersnote.com/v1/handlers"
	"gamersnote.com/v1/middlewares"
	"gamersnote.com/v1/restapi/operations"
)

//go:generate swagger generate server --target ../../api --name GamersnoteAPI --spec ../../swagger/swagger.yml --user *handlers.TokenPayload

func configureFlags(api *operations.GamersnoteAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GamersnoteAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	api.BearerAuth = utils.AuthHandler

	// users
	api.UserPostUserHandler = user.PostUserHandlerFunc(handlers.GetUsersHandler().PostUser)
	api.UserPutUserHandler = user.PutUserHandlerFunc(handlers.GetUsersHandler().PutMe)
	api.UserGetUserHandler = user.GetUserHandlerFunc(handlers.GetUsersHandler().GetUser)
	api.UserGetMeHandler = user.GetMeHandlerFunc(handlers.GetUsersHandler().GetMe)
	api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(handlers.GetUsersHandler().DeleteMe)

	// articles
	api.ArticlePostArticleHandler = article.PostArticleHandlerFunc(handlers.GetArticlesHandler().PostArticle)
	api.ArticleGetArticlesHandler = article.GetArticlesHandlerFunc(handlers.GetArticlesHandler().GetAllArticles)

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
func configureServer(s *http.Server, scheme, addr string) {
	configs.GetDB()
	configs.GetFirebaseApp()
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return middlewares.AccessLog(handler)
}
