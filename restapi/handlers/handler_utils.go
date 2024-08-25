package handers

import (
	"context"

	"github.com/john-k-antony/richpanel-blogposts-golang/constants"
	"github.com/john-k-antony/richpanel-blogposts-golang/models"
)

// GetHandlerContext create a context for invoking services from handlers. Returned context contains the user principal invoing this request
func GetHandlerContext(principal *models.Principal) context.Context {
	ctx := context.Background()
	handlerCtx := context.WithValue(ctx, constants.ContextKeyPrincipal, principal)
	return handlerCtx
}
