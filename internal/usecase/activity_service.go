package usecase

import (
	"context"
	"fmt"

	entity "github.com/Pancreasz/BackMor_Backend2/internal/entity"
	repository "github.com/Pancreasz/BackMor_Backend2/internal/interface/persistance"
	"github.com/google/uuid"
)

type ActivityService interface {
	ListActivities(ctx context.Context) ([]entity.Activity, error)
	CreateActivity(ctx context.Context, a entity.Activity) (entity.Activity, error)
	ListActivityMembers(ctx context.Context, activityID uuid.UUID) ([]entity.ActivityMemberResponse, error)
	JoinActivity(ctx context.Context, activityID, userID uuid.UUID) (entity.ActivityMemberResponse, error)
	GetActivityByID(ctx context.Context, id uuid.UUID) (entity.Activity, error)
	ListActivitiesByUser(ctx context.Context, userID uuid.UUID) ([]entity.Activity, error)
}

type activityService struct {
	repo repository.ActivityRepository
}

func NewActivityService(repo repository.ActivityRepository) ActivityService {
	return &activityService{repo: repo}
}

func (s *activityService) ListActivities(ctx context.Context) ([]entity.Activity, error) {
	return s.repo.List(ctx)
}

func (s *activityService) CreateActivity(ctx context.Context, a entity.Activity) (entity.Activity, error) {
	return s.repo.Create(ctx, a)
}

func (s *activityService) ListActivityMembers(ctx context.Context, activityID uuid.UUID) ([]entity.ActivityMemberResponse, error) {
	return s.repo.ListMembers(ctx, activityID)
}

func (s *activityService) JoinActivity(ctx context.Context, activityID, userID uuid.UUID) (entity.ActivityMemberResponse, error) {
	// Optional: check if already joined
	members, _ := s.repo.ListMembers(ctx, activityID)
	for _, m := range members {
		if m.UserID == userID {
			return entity.ActivityMemberResponse{}, fmt.Errorf("user already joined")
		}
	}
	return s.repo.AddMember(ctx, activityID, userID, "member")
}

func (s *activityService) GetActivityByID(ctx context.Context, id uuid.UUID) (entity.Activity, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *activityService) ListActivitiesByUser(ctx context.Context, userID uuid.UUID) ([]entity.Activity, error) {
	return s.repo.ListActivitiesByUser(ctx, userID)
}
