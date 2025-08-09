package service

import (
	entity "api-go/internal/domain"
	"context"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int32) (entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
}
