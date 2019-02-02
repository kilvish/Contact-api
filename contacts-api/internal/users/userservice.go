package users

import (
	"context"
	"time"

	"github.com/anil/contacts-api/internal/users/uuid"
)

type userService struct {
	repo Repo
}

// NewUserService returns new User Service
func NewUserService(repo Repo) Service {
	return &userService{
		repo: repo,
	}
}

func (us *userService) addUser(ctx context.Context, req *addUserRequest) (*responseData, error) {
	if err := req.validate(); err != nil {
		return nil, err
	}
	userID := uuid.GetUUID()
	userData := &User{
		Created:  time.Now(),
		Username: req.Username,
		Email:    req.Email,
		UserID:   userID,
	}
	if err := us.repo.AddUser(ctx, userData); err != nil {
		return nil, err
	}
	res := new(responseData)
	res.User = userData
	return res, nil
}

func (us *userService) getUser(ctx context.Context, req *getUserRequest) (*responseData, error) {
	err := req.validate()
	if err != nil {
		return nil, err
	}

	var userData *User
	if req.ID != nil {
		userData, err = us.repo.GetUserByID(ctx, *req.ID)
		if err != nil {
			return nil, err
		}
	} else {
		userData, err = us.repo.GetUserByEmail(ctx, *req.Email)
		if err != nil {
			return nil, err
		}
	}

	res := new(responseData)
	res.User = userData
	return res, nil
}
