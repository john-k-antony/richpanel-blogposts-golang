package blogpostservice

import (
	"context"

	"github.com/john-k-antony/richpanel-blogposts-golang/restapi/operations/blog_post"

	"github.com/john-k-antony/richpanel-blogposts-golang/models/apimodels"
	repository "github.com/john-k-antony/richpanel-blogposts-golang/repository"
)

func ListBlogposts(ctx context.Context, params blog_post.ListBlogpostsParams) ([]*apimodels.Blogpost, int, error) {

	offset := -1
	limit := -1

	if params.Offset != nil {
		offset = int(*params.Offset)
	}

	if params.Limit != nil {
		limit = int(*params.Limit)
	}

	blogPostModels := repository.List(offset, limit)
	return blogPostModels, repository.Size(), nil
}
