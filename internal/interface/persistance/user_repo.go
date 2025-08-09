package repository

import (
	entity "api-go/internal/domain"
	"api-go/pkg/database/user_database" // sqlc generated package
	"context"
	"database/sql"
)

type userRepository struct {
	queries *user_database.Queries
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		queries: user_database.New(db),
	}
}

func (r *userRepository) GetByID(ctx context.Context, id int32) (entity.User, error) {
	userTable, err := r.queries.GetUser(ctx, id)
	if err != nil {
		return entity.User{}, err
	}
	return mapUserTableToEntity(userTable), nil
}

func (r *userRepository) List(ctx context.Context) ([]entity.User, error) {
	userTables, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]entity.User, len(userTables))
	for i, u := range userTables {
		users[i] = mapUserTableToEntity(u)
	}
	return users, nil
}

func mapUserTableToEntity(u user_database.UserTable) entity.User {
	return entity.User{
		ID:   uint(u.ID),
		Name: u.Name,
		Sex:  u.Sex,
	}
}
