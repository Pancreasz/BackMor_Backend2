package repository

import (
	"context"
	"database/sql"
	"fmt"

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

func (r *userRepository) InsertUser(ctx context.Context, name string, sex string) (entity.User, error) {
	// Prepare params struct expected by sqlc generated method
	params := user_database.InsertUserParams{
		Name: name,
		Sex:  sex,
	}
	fmt.Println(params)
	// Call the generated InsertUser method
	user, err := r.queries.InsertUser(ctx, params)
	if err != nil {
		return entity.User{}, err
	}

	return mapUserTableToEntity(user), nil
}

func mapUserTableToEntity(u user_database.UserTable) entity.User {
	return entity.User{
		ID:   uint(u.ID),
		Name: u.Name,
		Sex:  u.Sex,
	}
}
