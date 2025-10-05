package entity

import (
	"time"

	"github.com/google/uuid"
)

type Activity struct {
	ID              uuid.UUID
	CreatorID       uuid.UUID
	Title           string
	Description     *string
	StartTime       time.Time
	EndTime         *time.Time
	MaxParticipants *int32
	Visibility      string
	Latitude        float64
	Longitude       float64
	Location        *string
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
