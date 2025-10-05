package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"

	entity "github.com/Pancreasz/BackMor_Backend2/internal/entity"
	"github.com/Pancreasz/BackMor_Backend2/pkg/database/user_database" // sqlc generated package
)

type userRepository struct {
	queries *user_database.Queries
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		queries: user_database.New(db),
	}
}

func (r *userRepository) GetByID(ctx context.Context, id uuid.UUID) (entity.User, error) {
	userRow, err := r.queries.GetUser(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return mapUserRowToEntity(userRow), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	userTable, err := r.queries.GetUserbyEmail(ctx, email)
	if err != nil {
		return entity.User{}, err
	}
	return mapUserTableToEntity(userTable), nil
}

func (r *userRepository) List(ctx context.Context) ([]entity.User, error) {
	rows, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]entity.User, len(rows))
	for i, u := range rows {
		users[i] = mapUserRowToEntity(u)
	}
	return users, nil
}

func (r *userRepository) InsertUser(ctx context.Context, email, passwordHash, displayName string, avatarURL, bio *string) (entity.User, error) {
	params := user_database.InsertUserParams{
		Email:        email,
		PasswordHash: passwordHash,
		DisplayName:  displayName,
		AvatarUrl:    avatarURL,
		Bio:          bio,
	}
	fmt.Println(params)
	userRow, err := r.queries.InsertUser(ctx, params)
	if err != nil {
		return entity.User{}, err
	}
	return mapUserRowToEntity(userRow), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	row, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:          row.ID,
		Email:       row.Email,
		DisplayName: row.DisplayName,
		AvatarURL:   row.AvatarUrl,
		Bio:         row.Bio,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
	}, nil
}

func mapUserRowToEntity(u user_database.User) entity.User {
	return entity.User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		DisplayName:  u.DisplayName,
		AvatarURL:    u.AvatarUrl,
		Bio:          u.Bio,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
	}
}
