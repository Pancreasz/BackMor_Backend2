package usecase

import (
	"errors"
)

var (
	ErrUserNotFound          = errors.New("user not found")
	ErrFailedToRetrieveUsers = errors.New("failed to retrieve the user")
	ErrFailedToInsertUser    = errors.New("failed to insert the user")
	ErrFailedToUpdateUser    = errors.New("Fail to update user infomation")
)
