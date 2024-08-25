package blogpostservice

import (
	"context"

	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations/blog_post"

	"github.com/john-k-antony/richpanel-blogposts-golang/models/apimodels"
	repository "github.com/john-k-antony/richpanel-blogposts-golang/repository"
)

func GetBlogpostByID(ctx context.Context, params blog_post.GetBlogpostByIDParams) (*apimodels.Blogpost, error) {
	blogPostModel, err := repository.Get(params.ID)
	if err != nil {
		return nil, nil
	} else {
		return blogPostModel, nil
	}
}
