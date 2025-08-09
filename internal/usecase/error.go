package usecase

import (
	"errors"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrFailedToRetrieveUsers = errors.New("failed to retrieve the user")
)
