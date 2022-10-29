package utils

import (
	"fmt"
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

func RespondError(c *gin.Context, statusCode int, msg string) {
	data := gin.H{"msg": msg}
	c.Error(fmt.Errorf(msg))
	ComposeResponse(false, statusCode, c, data)
}

func RespondInternalServerError(c *gin.Context) {
	RespondError(c, http.StatusInternalServerError, errors.ErrInternalServerError.Error())
}
