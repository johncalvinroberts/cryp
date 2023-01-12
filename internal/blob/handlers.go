package blob

import (
	"net/http"

	"github.com/johncalvinroberts/cryp/internal/utils"
	"github.com/johncalvinroberts/cryp/internal/whoami"
	"github.com/labstack/echo/v4"
)

func (svc *BlobService) HandleInitializeUpload(c echo.Context) error {
	file, err := c.FormFile("file")
	if err != nil {
		return utils.RespondError(c, http.StatusBadRequest, err)
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	claims := whoami.GetUserFromContext(c)
	email := claims.Email
	svc.UploadFile(src, email)
	return nil
}
