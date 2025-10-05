package usecase

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"

	entity "github.com/Pancreasz/BackMor_Backend2/internal/entity"
)

type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (entity.User, error)
	List(ctx context.Context) ([]entity.User, error)
	InsertUser(ctx context.Context, email string, passwordHash string, displayName string, avatarURL *string, bio *string) (entity.User, error)
	GetByEmail(ctx context.Context, email string) (entity.User, error)
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
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

func (s *UserService) InsertNewUser(ctx context.Context, email string, passwordHash string, displayName string, avatarURL *string, bio *string) (*entity.User, error) {
	user, err := s.repo.InsertUser(ctx, email, passwordHash, displayName, avatarURL, bio)
	fmt.Println("service", email, passwordHash, displayName, avatarURL, bio)
	if err != nil {
		return nil, ErrFailedToInsertUser
	}
	return &user, nil
}

func (s *UserService) GetUserByEmail(ctx context.Context, email string) (entity.User, error) {
	return s.repo.GetByEmail(ctx, email)
}
