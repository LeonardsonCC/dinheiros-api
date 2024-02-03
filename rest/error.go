package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Error struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

func Err(c *gin.Context, message string, err error) {
	c.JSON(http.StatusInternalServerError, Error{
		Message: message,
		Error:   err.Error(),
	})
}
