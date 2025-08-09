package http

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	response "github.com/Pancreasz/BackMor_Backend2/infrastructure/router"
	"github.com/Pancreasz/BackMor_Backend2/internal/entity"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserByID(ctx context.Context, id int32) (*entity.User, error)
	ListUsers(ctx context.Context) ([]entity.User, error)
}

type UserHandler struct {
	service UserService
}

func NewUserServiceHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	ctx := context.Background()
	userID, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.GetUserByID(ctx, int32(userID))
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
