package blogposthandlers

import (
	"net/http"
	"strconv"

	"github.com/go-openapi/runtime/middleware"
	"github.com/john-k-antony/richpanel-blogposts-golang/constants"
	"github.com/john-k-antony/richpanel-blogposts-golang/models"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/errorcode"
	handlers "github.com/john-k-antony/richpanel-blogposts-golang/restapi/handlers"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations/blog_post"
	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/responder"
	blogpostservice "github.com/john-k-antony/richpanel-blogposts-golang/services/blogpost"
)

func ListBlogpostsHandler(params blog_post.ListBlogpostsParams, principal *models.Principal) middleware.Responder {
	resp := responder.New(params.HTTPRequest)
	ctx := handlers.GetHandlerContext(principal)
	blogposts, totalSize, err := blogpostservice.ListBlogposts(ctx, params)
	if err != nil {
		return resp.Error(http.StatusUnprocessableEntity, errorcode.BlogPostGetError, err)
	}
	resp.Header(constants.ListResultsTotalCountHeader, strconv.FormatInt(int64(totalSize), 10))
	resp.Header(constants.ListResultsPageSizeCountHeader, strconv.FormatInt(int64(len(blogposts)), 10))
	return resp.OK(blogposts)
}
