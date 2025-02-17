package errors

import "errors"

var (
	//controllers
	ErrAuthRequiredError  = errors.New("auth is required")
	ErrRegistrationsError = errors.New("registrations are required")
	ErrAuthenticated      = errors.New("authentication is required for this action")
	ErrDataBindError      = errors.New("wrong data format")

	//repositories
	ErrEntityNotFound = errors.New("entity not found")

	//usecases
	ErrEntityAlreadyExists = errors.New("entity already exists")
	ErrPasswordMismatch    = errors.New("password mismatch")
)
