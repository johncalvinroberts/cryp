package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johncalvinroberts/cryp/internal/errors"
)

func ComposeResponse(success bool, statusCode int, c *gin.Context, data any) {
	c.JSON(statusCode, gin.H{
		"success": success,
		"data":    data,
	})
}

func RespondOK(c *gin.Context, data any) {
	ComposeResponse(true, http.StatusOK, c, data)
}

func RespondCreated(c *gin.Context, data any) {
	ComposeResponse(true, http.StatusCreated, c, data)
}

func RespondError(c *gin.Context, statusCode int, err error) {
	data := gin.H{"msg": err.Error()}
	c.Error(err)
	ComposeResponse(false, statusCode, c, data)
}

func RespondInternalServerError(c *gin.Context) {
	RespondError(c, http.StatusInternalServerError, errors.ErrInternalServerError)
}

func RespondUnauthorized(c *gin.Context, err error) {
	RespondError(c, http.StatusUnauthorized, errors.ErrInternalServerError)
}
