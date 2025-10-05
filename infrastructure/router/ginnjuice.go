package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendSuccessResponse[T any](c *gin.Context, data T) {
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
