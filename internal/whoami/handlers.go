package whoami

import (
	"net/http"

	"github.com/johncalvinroberts/cryp/internal/utils"

	"github.com/gin-gonic/gin"
)

func (svc *WhoamiService) HandleStartWhoamiChallenge(c *gin.Context) {
	req := &StartWhoamiChallengeDTO{}
	err := c.BindJSON(req)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
		return
	}
	err = svc.StartWhoamiChallenge(req.Email)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
		return
	}
	utils.RespondCreated(c, nil)
}

func (svc *WhoamiService) HandleTryWhoamiChallenge(c *gin.Context) {
	req := &TryWhoamiChallengeRequestDTO{}
	err := c.BindJSON(req)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
		return
	}
	jwt, err := svc.TryWhoamiChallenge(req.Email, req.OTP)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
		return
	}
	utils.RespondOK(c, &TryWhoamiChallengeResponseDTO{JWT: jwt})
}

func (svc *WhoamiService) HandleGetWhoami(c *gin.Context) {
	utils.RespondOK(c, "good")
}
