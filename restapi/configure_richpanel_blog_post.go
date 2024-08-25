// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	blogposthandlers "github.com/john-k-antony/richpanel-blogposts-golang/restapi/handlers/blogpost"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"

	"github.com/john-k-antony/richpanel-blogposts-golang/models"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations/blog_post"
)

var bearerRegEx = regexp.MustCompile("(?i)Bearer ")

//go:generate swagger generate server --target ../../richpanel-blogposts-golang --name RichpanelBlogPost --spec ../swagger/swagger-flat.tmp.yaml --model-package models/apimodels --principal models.Principal

func configureFlags(api *operations.RichpanelBlogPostAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RichpanelBlogPostAPI) http.Handler {
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
	api.APIKeyAuth = func(token string) (*models.Principal, error) {
		tokens := bearerRegEx.Split(token, -1)
		if len(tokens) > 1 {
			bearerToken := strings.TrimSpace(tokens[1])

			if bearerToken != "N2E5OTNmNzg3NTQ2MmM4YmFkNmNmODE3NDE4MGZkY2Q=" {
				return nil, errors.New(http.StatusUnauthorized, "Invalid Token")
			}

			userID := fmt.Sprintf("U-%s", bearerToken)
			// prin := models.Principal(bearerToken)
			prin := models.Principal{
				AccessToken: bearerToken,
				UserID:      userID,
			}
			return &prin, nil
		}
		return nil, errors.New(http.StatusUnauthorized, "Invalid Token")
	}

	api.BlogPostCreateBlogpostHandler = blog_post.CreateBlogpostHandlerFunc(blogposthandlers.CreateBlogpostHandler)
	api.BlogPostDeleteBlogpostByIDHandler = blog_post.DeleteBlogpostByIDHandlerFunc(blogposthandlers.DeleteBlogpostByIDHandler)
	api.BlogPostGetBlogpostByIDHandler = blog_post.GetBlogpostByIDHandlerFunc(blogposthandlers.GetBlogpostByIDHandler)
	api.BlogPostListBlogpostsHandler = blog_post.ListBlogpostsHandlerFunc(blogposthandlers.ListBlogpostsHandler)
	api.BlogPostUpdateBlogpostByIDHandler = blog_post.UpdateBlogpostByIDHandlerFunc(blogposthandlers.UpdateBlogpostByIDHandler)

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
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
