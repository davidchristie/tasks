package graphql

import "errors"

var ErrMustBeLoggedIn = errors.New("must be logged in")
var ErrNotFound = errors.New("not found")
