package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"

	response "github.com/Pancreasz/BackMor_Backend2/infrastructure/response"
	"github.com/Pancreasz/BackMor_Backend2/internal/entity"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	ListUsers(ctx context.Context) ([]entity.User, error)
	InsertNewUser(ctx context.Context, email string, passwordHash string, displayName string, avatarURL *string, bio *string) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (entity.User, error)
}

type UserHandler struct {
	service UserService
}

func NewUserServiceHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	ctx := c.Request.Context()

	userID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, fmt.Errorf("invalid UUID format"))
		return
	}

	user, err := h.service.GetUserByID(ctx, userID)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			response.SendErrorResponse(c, http.StatusNotFound, err)
			return
		}
		response.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.SendSuccessResponse(c, user)
}

func (h *UserHandler) GetAllUser(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := h.service.ListUsers(ctx)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}
	response.SendSuccessResponse(c, users)
}

func (h *UserHandler) InsertNewUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		Email        string  `json:"email" binding:"required,email"`
		PasswordHash string  `json:"password_hash" binding:"required"`
		DisplayName  string  `json:"display_name" binding:"required"`
		AvatarURL    *string `json:"avatar_url"`
		Bio          *string `json:"bio"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.InsertNewUser(ctx, req.Email, req.PasswordHash, req.DisplayName, req.AvatarURL, req.Bio)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.SendSuccessResponse(c, user)
}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	user, err := h.service.GetUserByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}
