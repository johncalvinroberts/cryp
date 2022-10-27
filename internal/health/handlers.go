package health

import (
	"github.com/gin-gonic/gin"
	"github.com/johncalvinroberts/cryp/internal/utils"
)

func HandleGetHealth(c *gin.Context) {
	healthy := GetHealth()
	if healthy {
		utils.RespondOK(c, nil)
	} else {
		utils.RespondInternalServerError(c)
	}
}
