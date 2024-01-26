package oauth2

import "errors"

// ErrUnauthorizedRequest user unauthorized error
var ErrUnauthorizedRequest = errors.New("unauthorized request")

// ErrInsufficientScope use does not have sufficient scopes for the requested operation
var ErrInsufficientScope = errors.New("insufficient scope")
