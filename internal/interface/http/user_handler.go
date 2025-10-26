package http

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	// "time"

	"github.com/google/uuid"

	// "github.com/Pancreasz/BackMor_Backend2/infrastructure/config"
	response "github.com/Pancreasz/BackMor_Backend2/infrastructure/response"
	"github.com/Pancreasz/BackMor_Backend2/internal/entity"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	ListUsers(ctx context.Context) ([]entity.User, error)
	InsertNewUser(ctx context.Context, email string, passwordHash string, displayName string, avatarURL *string, bio *string, sex *string, age *int) (*entity.User, error)
	UpdateUserProfile(ctx context.Context, displayName string, avatarURL *string, bio *string, sex *string, age *int, email string) (*entity.User, error)
	UpdateUserAvatarData(ctx context.Context, avatarData []byte, email string) (*entity.User, error)
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
		Sex          *string `json:"sex"` // optional
		Age          *int    `json:"age"` // optional
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.InsertNewUser(ctx, req.Email, req.PasswordHash, req.DisplayName, req.AvatarURL, req.Bio, req.Sex, req.Age)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.SendSuccessResponse(c, user)
}

func (h *UserHandler) UpdateUserProfile(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		Email       string  `json:"email" binding:"required,email"`
		DisplayName string  `json:"display_name" binding:"required"`
		AvatarURL   *string `json:"avatar_url"`
		Bio         *string `json:"bio"`
		Sex         *string `json:"sex"`
		Age         *int    `json:"age"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.UpdateUserProfile(ctx, req.DisplayName, req.AvatarURL, req.Bio, req.Sex, req.Age, req.Email)
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

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
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

func (h *UserHandler) UpdateUserAvatarData(c *gin.Context) {
	// Get file from form-data
	file, _, err := c.Request.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No file found"})
		return
	}
	defer file.Close()

	// Read file bytes
	avatarBytes, err := io.ReadAll(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read file"})
		return
	}

	// Get email from form-data
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
		return
	}

	// Update user avatar data in database
	updatedUser, err := h.service.UpdateUserAvatarData(c.Request.Context(), avatarBytes, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update avatar in database"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Avatar updated successfully",
		"url":     updatedUser.AvatarURL,
	})
}

// func (h *UserHandler) UpdateUserAvatarURL(c *gin.Context) {
// 	// Get file from form-data
// 	file, header, err := c.Request.FormFile("image")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "No file found"})
// 		return
// 	}
// 	defer file.Close()

// 	// Get email from form-data
// 	email := c.PostForm("email")
// 	if email == "" {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "email is required"})
// 		return
// 	}

// Generate unique filename
// objectName := fmt.Sprintf("uploads/%d_%s", time.Now().Unix(), header.Filename)

// // Upload to Firebase Storage
// ctx := context.Background()
// bucket, err := config.StorageClient.Bucket(config.BucketName)
// if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get bucket"})
// 	return
// }

// wc := bucket.Object(objectName).NewWriter(ctx)
// if _, err = io.Copy(wc, file); err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload"})
// 	return
// }

// if err := wc.Close(); err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to finalize upload: %v", err)})
// 	return
// }

// Construct public URL
// imageURL := fmt.Sprintf("https://firebasestorage.googleapis.com/v0/b/%s/o/%s?alt=media", config.BucketName, objectName)

// Update user avatar in database
// updatedUser, err := h.service.UpdateUserAvatar(ctx, &imageURL, email)
// if err != nil {
// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update avatar in database"})
// 	return
// }

// c.JSON(http.StatusOK, gin.H{
// 	"message": "Avatar updated successfully",
// 	"url":     updatedUser.AvatarURL,
// })
// }
