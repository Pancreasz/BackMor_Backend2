package response

import (
	"net/http"

	"github.com/Pancreasz/BackMor_Backend2/internal/entity"
	"github.com/gin-gonic/gin"
)

type EntityType interface {
	*entity.User | []entity.User
}

func SendSuccessResponse[T EntityType](c *gin.Context, data T) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   data,
	})
}

func SendErrorResponse(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": err.Error(),
	})
}
