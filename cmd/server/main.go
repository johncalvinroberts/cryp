package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/johncalvinroberts/cryp/internal/config"
	"github.com/johncalvinroberts/cryp/internal/email"
	"github.com/johncalvinroberts/cryp/internal/storage"

	"github.com/johncalvinroberts/cryp/internal/blob"
	"github.com/johncalvinroberts/cryp/internal/health"
	"github.com/johncalvinroberts/cryp/internal/whoami"
	"github.com/johncalvinroberts/cryp/ui"
)

func main() {
	config := config.InitAppConfig()
	e := echo.New()
	e.Logger.Print("Starting Server")
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RateLimiter(middleware.NewRateLimiterMemoryStore(20)))
	storageSrv := storage.InitStorageService(config.AWSSession, config.Timeout)
	emailSrv := email.InitEmailService(config)
	whoamiSrv := whoami.InitWhoamiService(config.JWTSecret, config.Storage.WhoamiBucketName, storageSrv, emailSrv)
	blobSrv := blob.InitBlobService(storageSrv, config.Storage.BlobBucketName, config.Storage.BlobPointerBucketName)
	e.GET("/", echo.WrapHandler(ui.GetHandler()))
	e.GET("/api/health", health.HandleGetHealth)
	// // whoami
	e.POST("/api/whoami/start", whoamiSrv.HandleStartWhoamiChallenge)
	e.POST("/api/whoami/try", whoamiSrv.HandleTryWhoamiChallenge)
	e.GET("/api/whoami", whoamiSrv.VerifyWhoamiMiddleware(whoamiSrv.HandleGetWhoami))
	e.POST("/api/whoami/refresh", whoamiSrv.VerifyWhoamiMiddleware(whoamiSrv.HandleRefreshWhoamiToken))
	e.DELETE("/api/whoami", whoamiSrv.VerifyWhoamiMiddleware(whoamiSrv.HandleDestroyWhoamiToken))
	e.DELETE("/api/whoami/everything", whoamiSrv.VerifyWhoamiMiddleware(whoamiSrv.HandleDestroyEverything))
	// uploads
	e.POST("/api/blobs", whoamiSrv.VerifyWhoamiMiddleware(blobSrv.HandleInitializeUpload))
	e.Logger.Fatal(e.Start("localhost:" + config.Port))
}
