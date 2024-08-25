package blogpostservice

import (
	"context"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations/blog_post"

	"github.com/john-k-antony/richpanel-blogposts-golang/models/apimodels"
	repository "github.com/john-k-antony/richpanel-blogposts-golang/repository"
)

func UpdateBlogpostByID(ctx context.Context, params blog_post.UpdateBlogpostByIDParams) (*apimodels.Blogpost, error) {
	blogPostModel, err := repository.Get(params.ID)

	if err != nil {
		return nil, nil
	}

	tNow := strfmt.DateTime(time.Now())

	if params.Body.Title != nil {
		blogPostModel.Title = params.Body.Title
	}

	if params.Body.Contents != nil {
		blogPostModel.Contents = *(params.Body.Contents)
	}
	blogPostModel.ModifiedAt = &tNow
	repository.Put(blogPostModel)

	return blogPostModel, nil
}
