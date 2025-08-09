package service

import (
	"context"

	entity "github.com/Pancreasz/BackMor_Backend2/internal/entity"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int32) (entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
}
