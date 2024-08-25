package blogposthandlers

import (
	"errors"
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/john-k-antony/richpanel-blogposts-golang/models"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/errorcode"
	handlers "github.com/john-k-antony/richpanel-blogposts-golang/restapi/handlers"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations/blog_post"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/responder"
	blogpostservice "github.com/john-k-antony/richpanel-blogposts-golang/services/blogpost"
)

func GetBlogpostByIDHandler(params blog_post.GetBlogpostByIDParams, principal *models.Principal) middleware.Responder {
	resp := responder.New(params.HTTPRequest)
	ctx := handlers.GetHandlerContext(principal)
	blogpost, err := blogpostservice.GetBlogpostByID(ctx, params)
	if err != nil {
		return resp.Error(http.StatusUnprocessableEntity, errorcode.BlogPostGetError, err)
	}
	if blogpost == nil {
		return resp.Error(http.StatusNotFound, errorcode.BlogPostGetError, errors.New("Blogpost Not Found"))
	}
	return resp.OK(blogpost)
}
