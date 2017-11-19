package datastore

import (
  "github.com/Typeform/users/transport/v1"
  "errors"
)

var (
  // ErrUserNotFound is returned when a DB can't find a user
  ErrUserNotFound = errors.New("user not found")
)

// Users represents a backing datastore for the Users service
type Users interface {
	GetUser(int) (*v1.User, error)
}
