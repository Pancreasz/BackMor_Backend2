package usecase

import (
	"context"
	"database/sql"
	"errors"

	entity "github.com/Pancreasz/BackMor_Backend2/internal/entity"
)

type UserRepository interface {
	GetByID(ctx context.Context, id int32) (entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(ctx context.Context, id int32) (*entity.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) { // sqlc returns this for no match
			return nil, ErrUserNotFound
		}
		return nil, ErrFailedToRetrieveUsers
	}
	return &user, nil
}

func (s *UserService) ListUsers(ctx context.Context) ([]entity.User, error) {
	users, err := s.repo.List(ctx)
	if err != nil {
		return nil, ErrFailedToRetrieveUsers
	}
	return users, nil
}
