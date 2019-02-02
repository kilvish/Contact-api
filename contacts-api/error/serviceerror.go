package error

import "errors"

// Errors used in the service
var (
	ErrUserNameMissingAddUserRequest = errors.New("User name missing in add user request")
	ErrEmailMissingAddUserRequest    = errors.New("Email missing for add user request")
	ErrInvalidGetUserRequest         = errors.New("Error invalid get users request, either email or user id required")
	ErrInvalidOrMissingBody          = errors.New("Invalid or Missing Reqeust Body, Add User request")
	ErrDecodingAddUserRequestBody    = errors.New("Error Decoding add user reqeust body")
	ErrInvalidRequest                = errors.New("Invalid request")
)

