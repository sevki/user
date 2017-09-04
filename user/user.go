package user

import (
	"context"
	"errors"

	"github.com/Typeform/users/transport/v1"
)

type userService struct {
}

func (us *userService) GetUser(ctx context.Context, id int) (*v1.User, error) {
	return nil, errors.New("not implemented")
}

// New retuns a new UserService as required by the transport layer
func New() v1.Users {
	return &userService{}
}
