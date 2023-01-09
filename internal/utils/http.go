package utils

import (
	"net/http"

	"github.com/johncalvinroberts/cryp/internal/errors"
	"github.com/labstack/echo/v4"
)

type CrypAPIResponse struct {
	Success bool `json:"success"`
	Data    any  `json:"data"`
}

func ComposeResponse(success bool, statusCode int, c echo.Context, data any) error {
	return c.JSON(statusCode, &CrypAPIResponse{Success: success, Data: data})
}

func RespondOK(c echo.Context, data any) error {
	return ComposeResponse(true, http.StatusOK, c, data)
}

func RespondCreated(c echo.Context, data any) error {
	return ComposeResponse(true, http.StatusCreated, c, data)
}

func RespondError(c echo.Context, statusCode int, err error) error {
	data := err.Error()
	return ComposeResponse(false, statusCode, c, data)
}

func RespondInternalServerError(c echo.Context) error {
	return RespondError(c, http.StatusInternalServerError, errors.ErrInternalServerError)
}

func RespondUnauthorized(c echo.Context, err error) error {
	return RespondError(c, http.StatusUnauthorized, err)
}
