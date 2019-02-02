package error

import "net/http"

var httpCodes = map[error]int{
	ErrUserNameMissingAddUserRequest: 400,
	ErrEmailMissingAddUserRequest:    400,
	ErrInvalidGetUserRequest:         400,
	ErrInvalidOrMissingBody:          400,
	ErrDecodingAddUserRequestBody:    400,
	ErrInvalidRequest:                400,
}

// GetHTTPCode gets status codes for errors
func GetHTTPCode(err error) int {
	code, ok := httpCodes[err]
	if !ok {
		return http.StatusInternalServerError
	}
	return code
}

