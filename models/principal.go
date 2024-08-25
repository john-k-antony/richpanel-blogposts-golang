package models

import (
	"github.com/go-openapi/strfmt"
)

// Principal principal
//
// keeps the Bearer access_token from the Authorization header
// type Principal string

// Principal principal
type Principal struct {
	AccessToken string
	UserName    string
	UserID      string
	UserRoles   []interface{}
}

// Validate validates this principal
func (m Principal) Validate(formats strfmt.Registry) error {
	return nil
}
