// Code generated by go-swagger; DO NOT EDIT.

package blog_post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/john-k-antony/richpanel-blogposts-golang/models"
)

// GetBlogpostByIDHandlerFunc turns a function with the right signature into a get blogpost by Id handler
type GetBlogpostByIDHandlerFunc func(GetBlogpostByIDParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn GetBlogpostByIDHandlerFunc) Handle(params GetBlogpostByIDParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// GetBlogpostByIDHandler interface for that can handle valid get blogpost by Id params
type GetBlogpostByIDHandler interface {
	Handle(GetBlogpostByIDParams, *models.Principal) middleware.Responder
}

// NewGetBlogpostByID creates a new http.Handler for the get blogpost by Id operation
func NewGetBlogpostByID(ctx *middleware.Context, handler GetBlogpostByIDHandler) *GetBlogpostByID {
	return &GetBlogpostByID{Context: ctx, Handler: handler}
}

/*
	GetBlogpostByID swagger:route GET /posts/{id} BlogPost getBlogpostById

# Get blog post by id

Retrieve the information of the blog post with the matching id.
*/
type GetBlogpostByID struct {
	Context *middleware.Context
	Handler GetBlogpostByIDHandler
}

func (o *GetBlogpostByID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetBlogpostByIDParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal *models.Principal
	if uprinc != nil {
		principal = uprinc.(*models.Principal) // this is really a models.Principal, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
