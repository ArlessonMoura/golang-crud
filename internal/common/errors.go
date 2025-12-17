package common

import "errors"

var (
	ErrUserNotFound   = errors.New("user not found")
	ErrInvalidInput   = errors.New("invalid input")
	ErrDuplicateEmail = errors.New("email already exists")
	ErrInternalServer = errors.New("internal server error")
)
