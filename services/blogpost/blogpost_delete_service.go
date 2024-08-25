package blogpostservice

import (
	"context"

	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations/blog_post"

	"github.com/john-k-antony/richpanel-blogposts-golang/models/apimodels"
	repository "github.com/john-k-antony/richpanel-blogposts-golang/repository"
)

func DeleteBlogpostByID(ctx context.Context, params blog_post.DeleteBlogpostByIDParams) (*apimodels.Blogpost, error) {
	blogPostModel, err := repository.Remove(params.ID)
	if err != nil {
		return nil, nil
	} else {
		return blogPostModel, nil
	}
}
