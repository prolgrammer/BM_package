package errors

import "errors"

var (
	ErrExpired = errors.New("token is expired")
	ErrInvalid = errors.New("token is invalid")
)
