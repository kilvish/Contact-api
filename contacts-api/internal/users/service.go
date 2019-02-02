package users

import "context"

// Service defines the propeties exposed by users service
type Service interface {
	addUser(context.Context, *addUserRequest) (*responseData, error)
	getUser(context.Context, *getUserRequest) (*responseData, error)
}

