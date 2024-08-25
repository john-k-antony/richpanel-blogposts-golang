package blogpostservice

import (
	"context"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/john-k-antony/richpanel-blogposts-golang/utils"

	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations/blog_post"

	"github.com/john-k-antony/richpanel-blogposts-golang/models/apimodels"
	repository "github.com/john-k-antony/richpanel-blogposts-golang/repository"
	services "github.com/john-k-antony/richpanel-blogposts-golang/services"
)

func CreateBlogpost(ctx context.Context, params blog_post.CreateBlogpostParams) (*apimodels.Blogpost, error) {
	var blogPostModel apimodels.Blogpost
	userPrincipal := services.GetUserPrincipalFromContext(ctx)
	tNow := strfmt.DateTime(time.Now())
	blogPostModel = apimodels.Blogpost{
		ID:        utils.ToStringP(utils.GetNextId()),
		Contents:  *params.Body.Contents,
		Title:     params.Body.Title,
		CreatedAt: &tNow,
		UserID:    utils.ToStringP(userPrincipal.UserID),
	}
	repository.Put(&blogPostModel)
	return &blogPostModel, nil
}
