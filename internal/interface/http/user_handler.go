package http

import (
	"context"
	"errors"
	"fmt"
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
	InsertNewUser(ctx context.Context, username string, name string, sex string, age int32, hashpass string, email string) (*entity.User, error)
}

type UserHandler struct {
	service UserService
}

func NewUserServiceHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) GetUserByID(c *gin.Context) {
	ctx := c.Request.Context()
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

func (h *UserHandler) InsertNewUser(c *gin.Context) {
	ctx := c.Request.Context()

	var req struct {
		Username string `json:"username" binding:"required"`
		Name     string `json:"name" binding:"required"`
		Sex      string `json:"sex" binding:"required"`
		Age      int32  `json:"age" binding:"required"`
		Hashpass string `json:"hashpass" binding:"required"`
		Email    string `json:"email" binding:"required"`
	}
	fmt.Println("handler", req.Name, req.Sex)
	if err := c.ShouldBindJSON(&req); err != nil {
		response.SendErrorResponse(c, http.StatusBadRequest, err)
		return
	}

	user, err := h.service.InsertNewUser(ctx, req.Username, req.Name, req.Sex, req.Age, req.Hashpass, req.Email)
	if err != nil {
		response.SendErrorResponse(c, http.StatusInternalServerError, err)
		return
	}

	response.SendSuccessResponse(c, user)
}
