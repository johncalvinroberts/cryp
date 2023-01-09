package health

import (
	"github.com/johncalvinroberts/cryp/internal/utils"
	"github.com/labstack/echo/v4"
)

func HandleGetHealth(c echo.Context) error {
	healthy := GetHealth()
	if healthy {
		return utils.RespondOK(c, nil)
	} else {
		return utils.RespondInternalServerError(c)
	}
}
