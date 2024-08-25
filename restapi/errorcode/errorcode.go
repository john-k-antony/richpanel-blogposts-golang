package errorcode

// internal error codes
const (
	InternalServerError = 1500
	AuthenticationError = 1401
	AuthorizationError  = 1403
)

const (
	BlogPostCreateError = 2001
	BlogPostUpdateError = 2002
	BlogPostGetError    = 2003
	BlogPostListError   = 2004
	BlogPostDeleteError = 2005
)
