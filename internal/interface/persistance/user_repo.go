package repository

import (
	"context"
	"database/sql"

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
	return mapGetUserRowToEntity(userRow), nil
}

func (r *userRepository) List(ctx context.Context) ([]entity.User, error) {
	rows, err := r.queries.ListUsers(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]entity.User, len(rows))
	for i, u := range rows {
		users[i] = mapListUsersRowToEntity(u)
	}
	return users, nil
}

func (r *userRepository) InsertUser(
	ctx context.Context,
	email, passwordHash, displayName string,
	avatarURL, bio *string,
	sex *string,
	age *int,
) (entity.User, error) {

	params := user_database.InsertUserParams{
		Email:        email,
		PasswordHash: passwordHash,
		DisplayName:  displayName,
		AvatarUrl:    avatarURL,
		Bio:          bio,
		Sex:          toNullString(sex),
		Age:          toNullInt32(age),
	}

	userRow, err := r.queries.InsertUser(ctx, params)
	if err != nil {
		return entity.User{}, err
	}

	return mapInsertUserRowToEntity(userRow), nil
}

func (r *userRepository) UpdateUserProfile(
	ctx context.Context,
	displayName string,
	avatarURL *string,
	bio *string,
	sex *string,
	age *int,
	email string,
) (entity.User, error) {

	params := user_database.UpdateUserProfileParams{
		DisplayName: displayName,
		AvatarUrl:   avatarURL,
		Bio:         bio,
		Sex:         toNullString(sex),
		Age:         toNullInt32(age),
		Email:       email,
	}

	userRow, err := r.queries.UpdateUserProfile(ctx, params)
	if err != nil {
		return entity.User{}, err
	}

	return mapUpdateUserProfileRowToEntity(userRow), nil
}

func (r *userRepository) GetByEmail(ctx context.Context, email string) (entity.User, error) {
	row, err := r.queries.GetUserbyEmail(ctx, email)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:          row.ID,
		Email:       row.Email,
		DisplayName: row.DisplayName,
		AvatarURL:   row.AvatarUrl,
		AvatarData:  row.AvatarData,
		Bio:         row.Bio,
		CreatedAt:   row.CreatedAt,
		UpdatedAt:   row.UpdatedAt,
		Sex:         nullStringToPtr(row.Sex),
		Age:         nullInt32ToPtr(row.Age),
	}, nil
}

func (r *userRepository) UpdateUserAvatarData(
	ctx context.Context,
	avatarData []byte,
	email string,
) (entity.User, error) {

	params := user_database.UpdateUserAvatarDataParams{
		AvatarData: avatarData,
		Email:      email,
	}

	userRow, err := r.queries.UpdateUserAvatarData(ctx, params)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:           userRow.ID,
		Email:        userRow.Email,
		PasswordHash: userRow.PasswordHash,
		DisplayName:  userRow.DisplayName,
		AvatarURL:    userRow.AvatarUrl,
		AvatarData:   userRow.AvatarData, // now we have the actual data
		Bio:          userRow.Bio,
		CreatedAt:    userRow.CreatedAt,
		UpdatedAt:    userRow.UpdatedAt,
		Sex:          nullStringToPtr(userRow.Sex),
		Age:          nullInt32ToPtr(userRow.Age),
	}, nil
}

func (r *userRepository) UpdateUserAvatarURL(
	ctx context.Context,
	email string,
	avatarURL string,
) (entity.User, error) {
	params := user_database.UpdateUserAvatarURLParams{
		AvatarUrl: &avatarURL,
		Email:     email,
	}

	userRow, err := r.queries.UpdateUserAvatarURL(ctx, params)
	if err != nil {
		return entity.User{}, err
	}

	return entity.User{
		ID:           userRow.ID,
		Email:        userRow.Email,
		PasswordHash: userRow.PasswordHash,
		DisplayName:  userRow.DisplayName,
		AvatarURL:    userRow.AvatarUrl,
		Bio:          userRow.Bio,
		CreatedAt:    userRow.CreatedAt,
		UpdatedAt:    userRow.UpdatedAt,
		Sex:          nullStringToPtr(userRow.Sex),
		Age:          nullInt32ToPtr(userRow.Age),
	}, nil
}

// -----------------
// Mappers for each sqlc row type
// -----------------

func mapGetUserRowToEntity(u user_database.GetUserRow) entity.User {
	return entity.User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		DisplayName:  u.DisplayName,
		AvatarURL:    u.AvatarUrl,
		AvatarData:   nil,
		Bio:          u.Bio,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		Sex:          nullStringToPtr(u.Sex),
		Age:          nullInt32ToPtr(u.Age),
	}
}

func mapListUsersRowToEntity(u user_database.ListUsersRow) entity.User {
	return entity.User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		DisplayName:  u.DisplayName,
		AvatarURL:    u.AvatarUrl,
		AvatarData:   nil,
		Bio:          u.Bio,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		Sex:          nullStringToPtr(u.Sex),
		Age:          nullInt32ToPtr(u.Age),
	}
}

func mapInsertUserRowToEntity(u user_database.InsertUserRow) entity.User {
	return entity.User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		DisplayName:  u.DisplayName,
		AvatarURL:    u.AvatarUrl,
		AvatarData:   nil,
		Bio:          u.Bio,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		Sex:          nullStringToPtr(u.Sex),
		Age:          nullInt32ToPtr(u.Age),
	}
}

func mapUpdateUserProfileRowToEntity(u user_database.UpdateUserProfileRow) entity.User {
	return entity.User{
		ID:           u.ID,
		Email:        u.Email,
		PasswordHash: u.PasswordHash,
		DisplayName:  u.DisplayName,
		AvatarURL:    u.AvatarUrl,
		AvatarData:   nil, // <-- include AvatarData here
		Bio:          u.Bio,
		CreatedAt:    u.CreatedAt,
		UpdatedAt:    u.UpdatedAt,
		Sex:          nullStringToPtr(u.Sex),
		Age:          nullInt32ToPtr(u.Age),
	}
}

// -----------------
// Helper functions
// -----------------

func nullStringToPtr(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func nullInt32ToPtr(ni sql.NullInt32) *int {
	if ni.Valid {
		i := int(ni.Int32)
		return &i
	}
	return nil
}

func toNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{String: *s, Valid: true}
	}
	return sql.NullString{Valid: false}
}

func toNullInt32(i *int) sql.NullInt32 {
	if i != nil {
		return sql.NullInt32{Int32: int32(*i), Valid: true}
	}
	return sql.NullInt32{Valid: false}
}
