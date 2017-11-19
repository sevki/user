package user

import (
	"context"

	"github.com/Typeform/users/datastore"
	"github.com/Typeform/users/transport/v1"
)

type userService struct {
	db datastore.Users
}

// New returns a new v1.Users as required by the transport layer, it accepts
// a datastore.Users as it's only argument.
func New(db datastore.Users) v1.Users {
	return &userService{
		db: db,
	}
}

func (us *userService) GetUser(ctx context.Context, id int) (*v1.User, error) {
	return us.db.GetUser(id)
}
