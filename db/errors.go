package db

import "errors"

var (
	// ErrUnknownAttribute returns when an invalid sorting attribute is selected
	ErrUnknownAttribute = errors.New("unknown attribute")
	// ErrUnknownType returns when an unknown type is received by the client
	ErrUnknownType = errors.New("unknown type")
	// ErrUnknownUser returns when an unknown user is requested by the client
	ErrUnknownUser = errors.New("unknown user")
)
