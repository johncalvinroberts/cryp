package whoami

import (
	"log"
	"net/http"

	"github.com/johncalvinroberts/cryp/internal/errors"
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
	// TODO: get info about current token holder
	// uploads, credits, etc.
	utils.RespondOK(c, "good")
}

func (svc *WhoamiService) HandleRefreshWhoamiToken(c *gin.Context) {
	token := svc.extractTokenFromRequest(c)
	claims, err := svc.tokenService.VerifyTokenAndParseClaims(token)
	// error parsing JWT
	if err != nil {
		utils.RespondUnauthorized(c, err)
		return
	}
	jwt, err := svc.RefreshWhoamiToken(token, claims)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
		return
	}
	utils.RespondOK(c, &RefreshWhoamiTokenResponseDTO{JWT: jwt})

}

func (svc *WhoamiService) HandleDestroyWhoamiToken(c *gin.Context) {
	token := svc.extractTokenFromRequest(c)
	if token == "" {
		// don't do anything
		log.Println("token gone. exiting.")
		return
	}
	err := svc.tokenService.EvictToken(token)
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, err)
	} else {
		utils.RespondOK(c, nil)
	}

}

func (svc *WhoamiService) HandleDestroyEverything(c *gin.Context) {
	utils.RespondError(c, http.StatusExpectationFailed, errors.ErrNotImplemented)
}
