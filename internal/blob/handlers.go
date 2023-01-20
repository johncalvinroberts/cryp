package blob

import (
	"net/http"

	"github.com/johncalvinroberts/cryp/internal/utils"
	"github.com/johncalvinroberts/cryp/internal/whoami"
	"github.com/labstack/echo/v4"
)

func (svc *BlobService) HandleCreateBlob(c echo.Context) error {

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
	blob, err := svc.UploadFile(src, email)
	// TODO: more granular error handling
	if err != nil {
		return utils.RespondError(c, http.StatusBadRequest, err)
	}
	res := &UploadBlobResponseDTO{
		Blob{
			Url:       blob.Url,
			CreatedAt: blob.CreatedAt,
			UpdatedAt: blob.UpdatedAt,
			Title:     blob.Title,
		},
	}
	return utils.RespondOK(c, res)
}

func (svc *BlobService) HandleListBlobs(c echo.Context) error {
	claims := whoami.GetUserFromContext(c)
	email := claims.Email
	ptr, err := svc.ListBlobs(email)
	if err != nil {
		return utils.RespondError(c, http.StatusBadRequest, err)
	}
	res := &BlobPointersResponseDTO{
		BlobPointers{
			Blobs: ptr.Blobs, Count: ptr.Count,
		},
	}
	return utils.RespondOK(c, res)
}
