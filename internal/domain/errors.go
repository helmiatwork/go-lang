package domain

import "errors"

var (
	ErrNotFound       = errors.New("user not found")
	ErrDuplicateEmail = errors.New("email already exists")
	ErrNameRequired   = errors.New("name is required")
	ErrInvalidEmail   = errors.New("invalid email")
)
