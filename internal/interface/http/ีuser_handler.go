package http

import (
	"context"
	"errors"
	"strconv"

	response "github.com/Pancreasz/BackMor_Backend2/infrastructure/router"
	"github.com/Pancreasz/BackMor_Backend2/internal/entity"
	"github.com/Pancreasz/BackMor_Backend2/internal/usecase"
	"github.com/gofiber/fiber/v2"
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

func (h *UserHandler) GetUserByID(c *fiber.Ctx) error {
	ctx := context.Background()
	user_id, err := strconv.ParseInt(c.Params("id"), 10, 32)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusBadRequest, err)
	}

	user, err := h.service.GetUserByID(ctx, int32(user_id))
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			return response.SendErrorResponse(c, fiber.StatusNotFound, err)
		}
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, user)
}

func (h *UserHandler) GetAllUser(c *fiber.Ctx) error {
	ctx := context.Background()
	users, err := h.service.ListUsers(ctx)
	if err != nil {
		return response.SendErrorResponse(c, fiber.StatusInternalServerError, err)
	}
	return response.SendSuccessResponse(c, users)
}
