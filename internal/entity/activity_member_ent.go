package entity

import (
	"time"

	"github.com/google/uuid"
)

type ActivityMemberResponse struct {
	UserID      uuid.UUID `json:"user_id"`
	DisplayName string    `json:"display_name"`
	Role        string    `json:"role"`
	JoinedAt    time.Time `json:"joined_at"`
}
