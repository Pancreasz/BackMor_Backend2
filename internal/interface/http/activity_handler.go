package http

import (
	"net/http"
	"time"

	response "github.com/Pancreasz/BackMor_Backend2/infrastructure/response"
	"github.com/Pancreasz/BackMor_Backend2/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/net/context"
)

type ActivityService interface {
	ListActivities(ctx context.Context) ([]entity.Activity, error)
	CreateActivity(ctx context.Context, a entity.Activity) (entity.Activity, error)
	ListActivityMembers(ctx context.Context, activityID uuid.UUID) ([]entity.ActivityMemberResponse, error)
	JoinActivity(ctx context.Context, activityID, userID uuid.UUID) (entity.ActivityMemberResponse, error)
	GetActivityByID(ctx context.Context, id uuid.UUID) (entity.Activity, error)
	ListActivitiesByUser(ctx context.Context, userID uuid.UUID) ([]entity.Activity, error)
	DeleteActivity(ctx context.Context, activityID uuid.UUID) error
	RemoveActivityMember(ctx context.Context, activityID, userID uuid.UUID) error
}

type ActivityHandler struct {
	service ActivityService
}

func NewActivityHandler(service ActivityService) *ActivityHandler {
	return &ActivityHandler{service: service}
}

func (h *ActivityHandler) ListActivities(c *gin.Context) {
	ctx := c.Request.Context()

	activities, err := h.service.ListActivities(ctx)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.SendSuccessResponse(c, activities)
}

func (h *ActivityHandler) GetActivityMembers(c *gin.Context) {
	idStr := c.Param("id")
	activityID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity id"})
		return
	}

	members, err := h.service.ListActivityMembers(c.Request.Context(), activityID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, members)
}

func (h *ActivityHandler) CreateActivity(c *gin.Context) {
	var req struct {
		CreatorID       uuid.UUID  `json:"creator_id" binding:"required"`
		Title           string     `json:"title" binding:"required"`
		Description     *string    `json:"description"`
		StartTime       time.Time  `json:"start_time" binding:"required"`
		EndTime         *time.Time `json:"end_time"`
		MaxParticipants *int32     `json:"max_participants"`
		Visibility      string     `json:"visibility" binding:"required,oneof=public friends"`
		Latitude        float64    `json:"latitude" binding:"required"`
		Longitude       float64    `json:"longitude" binding:"required"`
		Location        *string    `json:"location" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activity := entity.Activity{
		CreatorID:       req.CreatorID,
		Title:           req.Title,
		Description:     req.Description,
		StartTime:       req.StartTime,
		EndTime:         req.EndTime,
		MaxParticipants: req.MaxParticipants,
		Visibility:      req.Visibility,
		Latitude:        req.Latitude,
		Longitude:       req.Longitude,
		Location:        req.Location,
	}

	created, err := h.service.CreateActivity(c.Request.Context(), activity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, created)
}

func (h *ActivityHandler) JoinActivity(c *gin.Context) {
	var req struct {
		ActivityID uuid.UUID `json:"activity_id" binding:"required"`
		UserID     uuid.UUID `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	member, err := h.service.JoinActivity(c.Request.Context(), req.ActivityID, req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, member)
}

func (h *ActivityHandler) GetActivityByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity id"})
		return
	}

	activity, err := h.service.GetActivityByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "activity not found"})
		return
	}

	c.JSON(http.StatusOK, activity)
}

func (h *ActivityHandler) ListActivitiesByUser(c *gin.Context) {
	userID, err := uuid.Parse(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user id"})
		return
	}

	activities, err := h.service.ListActivitiesByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, activities)
}

func (h *ActivityHandler) DeleteActivity(c *gin.Context) {
	idStr := c.Param("id")
	activityID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid activity id"})
		return
	}

	if err := h.service.DeleteActivity(c.Request.Context(), activityID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}

func (h *ActivityHandler) RemoveActivityMember(c *gin.Context) {
	var req struct {
		ActivityID uuid.UUID `json:"activity_id" binding:"required"`
		UserID     uuid.UUID `json:"user_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.service.RemoveActivityMember(c.Request.Context(), req.ActivityID, req.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
