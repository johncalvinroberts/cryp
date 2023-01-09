package whoami

import (
	"net/http"

	"github.com/johncalvinroberts/cryp/internal/errors"
	"github.com/johncalvinroberts/cryp/internal/utils"
	"github.com/labstack/echo/v4"
)

func (svc *WhoamiService) HandleStartWhoamiChallenge(c echo.Context) error {
	req := &StartWhoamiChallengeDTO{}
	err := c.Bind(req)
	if err != nil {
		return utils.RespondError(c, http.StatusBadRequest, err)
	}
	err = svc.StartWhoamiChallenge(req.Email)
	if err != nil {
		return utils.RespondError(c, http.StatusBadRequest, err)
	}
	return utils.RespondCreated(c, nil)
}

func (svc *WhoamiService) HandleTryWhoamiChallenge(c echo.Context) error {
	req := &TryWhoamiChallengeRequestDTO{}
	err := c.Bind(req)
	if err != nil {
		return utils.RespondError(c, http.StatusBadRequest, err)
	}
	jwt, err := svc.TryWhoamiChallenge(req.Email, req.OTP)
	if err != nil {
		return utils.RespondError(c, http.StatusBadRequest, err)
	}
	return utils.RespondOK(c, &TryWhoamiChallengeResponseDTO{JWT: jwt})
}

func (svc *WhoamiService) HandleGetWhoami(c echo.Context) error {
	// TODO: get info about current token holder
	// uploads, credits, etc.
	return utils.RespondOK(c, "good")
}

func (svc *WhoamiService) HandleRefreshWhoamiToken(c echo.Context) error {
	token := svc.extractTokenFromRequest(c)
	claims, err := svc.tokenService.VerifyTokenAndParseClaims(token)
	// error parsing JWT
	if err != nil {
		return utils.RespondUnauthorized(c, err)
	}
	jwt, err := svc.RefreshWhoamiToken(token, claims)
	if err != nil {
		return utils.RespondError(c, http.StatusBadRequest, err)
	}
	return utils.RespondOK(c, &RefreshWhoamiTokenResponseDTO{JWT: jwt})
}

func (svc *WhoamiService) HandleDestroyWhoamiToken(c echo.Context) error {
	token := svc.extractTokenFromRequest(c)
	err := svc.tokenService.EvictToken(token)
	if err != nil {
		return utils.RespondError(c, http.StatusBadRequest, err)
	} else {
		return utils.RespondOK(c, nil)
	}
}

func (svc *WhoamiService) HandleDestroyEverything(c echo.Context) error {
	return utils.RespondError(c, http.StatusExpectationFailed, errors.ErrNotImplemented)
}

func (svc *WhoamiService) VerifyWhoamiMiddleware(endpointHandler echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := svc.extractTokenFromRequest(c)
		claims, err := svc.tokenService.VerifyTokenAndParseClaims(token)
		// error parsing JWT
		if err != nil {
			return utils.RespondUnauthorized(c, err)
		}
		c.Set(CTX_JWT_KEY, token)
		c.Set(CTX_JWT_CLAIMS_KEY, &claims)
		return endpointHandler(c)
	}
}

func (svc *WhoamiService) extractTokenFromRequest(c echo.Context) string {
	token := c.Request().Header.Get("Authorization")
	// no token in header
	if token == "" {
		utils.RespondUnauthorized(c, errors.ErrUnauthorized)
	}
	return token
}
