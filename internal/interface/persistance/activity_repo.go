package repository

import (
	"context"
	"database/sql"

	entity "github.com/Pancreasz/BackMor_Backend2/internal/entity"
	activity_database "github.com/Pancreasz/BackMor_Backend2/pkg/database/activity_database" // sqlc generated package
	"github.com/google/uuid"
)

type ActivityRepository interface {
	List(ctx context.Context) ([]entity.Activity, error)
	Create(ctx context.Context, a entity.Activity) (entity.Activity, error)
	ListMembers(ctx context.Context, activityID uuid.UUID) ([]entity.ActivityMemberResponse, error)
	AddMember(ctx context.Context, activityID, userID uuid.UUID, role string) (entity.ActivityMemberResponse, error)
	GetByID(ctx context.Context, id uuid.UUID) (entity.Activity, error)
	ListActivitiesByUser(ctx context.Context, userID uuid.UUID) ([]entity.Activity, error)
	Delete(ctx context.Context, activityID uuid.UUID) error
	RemoveMember(ctx context.Context, activityID, userID uuid.UUID) error
	// Later: Update, Delete
}

type activityRepository struct {
	queries *activity_database.Queries
}

func NewActivityRepository(db *sql.DB) ActivityRepository {
	return &activityRepository{
		queries: activity_database.New(db),
	}
}
func (r *activityRepository) List(ctx context.Context) ([]entity.Activity, error) {
	rows, err := r.queries.ListActivities(ctx)
	if err != nil {
		return nil, err
	}
	activities := make([]entity.Activity, len(rows))
	for i, a := range rows {
		activities[i] = mapActivity(a)
	}
	return activities, nil
}
func (r *activityRepository) Create(ctx context.Context, a entity.Activity) (entity.Activity, error) {
	params := activity_database.InsertActivityParams{
		CreatorID:       a.CreatorID,
		Title:           a.Title,
		Description:     a.Description,
		StartTime:       a.StartTime,
		EndTime:         a.EndTime,
		MaxParticipants: a.MaxParticipants,
		Visibility:      a.Visibility,
		Longitude:       a.Longitude,
		Latitude:        a.Latitude,
		Location:        a.Location,
	}
	row, err := r.queries.InsertActivity(ctx, params)
	if err != nil {
		return entity.Activity{}, err
	}
	return mapActivity(row), nil
}

func (r *activityRepository) ListMembers(ctx context.Context, activityID uuid.UUID) ([]entity.ActivityMemberResponse, error) {
	rows, err := r.queries.ListActivityMembers(ctx, activityID)
	if err != nil {
		return nil, err
	}

	members := make([]entity.ActivityMemberResponse, 0, len(rows))
	for _, row := range rows {
		members = append(members, entity.ActivityMemberResponse{
			UserID:      row.UserID,
			DisplayName: row.DisplayName,
			Role:        row.Role,
			JoinedAt:    row.JoinedAt,
		})
	}
	return members, nil
}
func (r *activityRepository) AddMember(ctx context.Context, activityID, userID uuid.UUID, role string) (entity.ActivityMemberResponse, error) {
	row, err := r.queries.AddActivityMember(ctx, activity_database.AddActivityMemberParams{
		ActivityID: activityID,
		UserID:     userID,
		Role:       role,
	})

	if err != nil {
		return entity.ActivityMemberResponse{}, err
	}
	return entity.ActivityMemberResponse{
		UserID: row.UserID,
		// DisplayName: "", // optional: join with users if you want name immediately
		Role:     row.Role,
		JoinedAt: row.JoinedAt,
	}, nil
}

func (r *activityRepository) GetByID(ctx context.Context, id uuid.UUID) (entity.Activity, error) {
	row, err := r.queries.GetActivityByID(ctx, id)
	if err != nil {
		return entity.Activity{}, err
	}
	return mapActivity(row), nil
}

func (r *activityRepository) ListActivitiesByUser(ctx context.Context, userID uuid.UUID) ([]entity.Activity, error) {
	rows, err := r.queries.ListActivitiesByUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	activities := make([]entity.Activity, 0, len(rows))
	for _, row := range rows {
		activities = append(activities, mapActivity(row))
	}
	return activities, nil
}

func (r *activityRepository) Delete(ctx context.Context, activityID uuid.UUID) error {
	return r.queries.DeleteActivity(ctx, activityID)
}

func (r *activityRepository) RemoveMember(ctx context.Context, activityID, userID uuid.UUID) error {
	return r.queries.DeleteActivityMember(ctx, activity_database.DeleteActivityMemberParams{
		ActivityID: activityID,
		UserID:     userID,
	})
}

// Helper to map sqlc struct → entity struct
func mapActivity(a activity_database.Activity) entity.Activity {
	return entity.Activity{
		ID:              a.ID,
		CreatorID:       a.CreatorID,
		Title:           a.Title,
		Description:     a.Description,
		StartTime:       a.StartTime,
		EndTime:         a.EndTime,
		MaxParticipants: a.MaxParticipants,
		Visibility:      a.Visibility,
		Latitude:        a.Latitude,
		Longitude:       a.Longitude,
		Location:        a.Location,
		CreatedAt:       a.CreatedAt,
		UpdatedAt:       a.UpdatedAt,
	}
}
