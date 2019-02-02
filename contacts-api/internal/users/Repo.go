package users

import "context"

type Repo interface {
	AddUser(ctx context.Context, user *User) error
	GetUserByID(ctx context.Context, ID string) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
}
