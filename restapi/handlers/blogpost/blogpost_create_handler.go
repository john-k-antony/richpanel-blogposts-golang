package blogposthandlers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
	"github.com/john-k-antony/richpanel-blogposts-golang/models"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/errorcode"
	handlers "github.com/john-k-antony/richpanel-blogposts-golang/restapi/handlers"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations/blog_post"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/responder"
	blogpostservice "github.com/john-k-antony/richpanel-blogposts-golang/services/blogpost"
)

func CreateBlogpostHandler(params blog_post.CreateBlogpostParams, principal *models.Principal) middleware.Responder {
	resp := responder.New(params.HTTPRequest)
	ctx := handlers.GetHandlerContext(principal)
	blogpost, err := blogpostservice.CreateBlogpost(ctx, params)
	if err != nil {
		return resp.Error(http.StatusUnprocessableEntity, errorcode.BlogPostGetError, err)
	}
	return resp.Created(blogpost)

}
