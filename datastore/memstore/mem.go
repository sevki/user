package memstore

import (
	"github.com/Typeform/users/datastore"
	"github.com/Typeform/users/transport/v1"
)

// Memdb is a inmemory users.DB implementation
type memdb struct {
	users map[int]v1.User
}

func (m *memdb) GetUser(id int) (*v1.User, error) {
	if user, ok := m.users[id]; ok {
		return &user, nil
	}
	return nil, datastore.ErrUserNotFound
}

// New returns an in-memory datastore.Users
func New(s map[int]v1.User) datastore.Users {
	return &memdb{
		users: s,
	}
}
