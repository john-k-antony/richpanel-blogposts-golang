// Code generated by go-swagger; DO NOT EDIT.

package blog_post

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	"github.com/john-k-antony/richpanel-blogposts-golang/models"
)

// ListBlogpostsHandlerFunc turns a function with the right signature into a list blogposts handler
type ListBlogpostsHandlerFunc func(ListBlogpostsParams, *models.Principal) middleware.Responder

// Handle executing the request and returning a response
func (fn ListBlogpostsHandlerFunc) Handle(params ListBlogpostsParams, principal *models.Principal) middleware.Responder {
	return fn(params, principal)
}

// ListBlogpostsHandler interface for that can handle valid list blogposts params
type ListBlogpostsHandler interface {
	Handle(ListBlogpostsParams, *models.Principal) middleware.Responder
}

// NewListBlogposts creates a new http.Handler for the list blogposts operation
func NewListBlogposts(ctx *middleware.Context, handler ListBlogpostsHandler) *ListBlogposts {
	return &ListBlogposts{Context: ctx, Handler: handler}
}

/*
	ListBlogposts swagger:route GET /posts BlogPost listBlogposts

# List blog posts, with pagination

Retrieve the information of the blog posts.
*/
type ListBlogposts struct {
	Context *middleware.Context
	Handler ListBlogpostsHandler
}

func (o *ListBlogposts) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewListBlogpostsParams()
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
