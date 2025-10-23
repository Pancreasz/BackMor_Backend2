package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID
	Email        string
	PasswordHash string
	DisplayName  string
	AvatarURL    *string
	Bio          *string
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Sex          *string
	Age          *int
}
