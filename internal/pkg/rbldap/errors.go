package rbldap

import "errors"

var (
	errNotImplemented = errors.New("dry-run not implemented")
	errUser404        = errors.New("User not found")
	errDuplicateUser  = errors.New("User Already exists")
)
