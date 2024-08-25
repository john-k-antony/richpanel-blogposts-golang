package authservice

import (
	"context"

	"github.com/john-k-antony/richpanel-blogposts-golang/constants"
	"github.com/john-k-antony/richpanel-blogposts-golang/models"
)

// GetUserPrincipalFromContext gets user principal in the context, if available
func GetUserPrincipalFromContext(ctx context.Context) *models.Principal {
	userPrincipal := ctx.Value(constants.ContextKeyPrincipal)
	if userPrincipal != nil {
		userPrincipalVal, ok := userPrincipal.(*models.Principal)
		if ok {
			return userPrincipalVal
		}
	}
	return nil
}
