package main

import (
	"log"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/johncalvinroberts/cryp/internal/config"
	"github.com/johncalvinroberts/cryp/internal/email"
	"github.com/johncalvinroberts/cryp/internal/health"
	"github.com/johncalvinroberts/cryp/internal/storage"
	"github.com/johncalvinroberts/cryp/internal/whoami"
	"github.com/johncalvinroberts/cryp/ui"
)

func main() {
	log.Print("Starting Server")
	config := config.InitAppConfig()
	gin.SetMode(config.GinMode)
	router := gin.Default()
	storageSrv := storage.InitStorageService(config.AWSSession, config.Timeout)
	emailSrv := email.InitEmailService(config)
	whoamiSrv := whoami.InitWhoamiService(config.JWTSecret, config.Storage.WhoamiBucketName, storageSrv, emailSrv)
	router.Use(static.Serve("/", ui.GetUIFileSystem()))
	router.GET("/api/health", health.HandleGetHealth)
	// whoami
	router.POST("/api/whoami/start", whoamiSrv.HandleStartWhoamiChallenge)
	router.POST("/api/whoami/try", whoamiSrv.HandleTryWhoamiChallenge)
	router.GET("/api/whoami", whoamiSrv.VerifyWhoamiMiddleware(whoamiSrv.HandleGetWhoami))
	router.POST("/api/whoami/refresh", whoamiSrv.VerifyWhoamiMiddleware(whoamiSrv.HandleRefreshWhoamiToken))
	router.DELETE("/api/whoami", whoamiSrv.VerifyWhoamiMiddleware(whoamiSrv.HandleDestroyWhoamiToken))
	router.DELETE("/api/whoami/everything", whoamiSrv.VerifyWhoamiMiddleware(whoamiSrv.HandleDestroyEverything))
	// credits
	log.Fatal(router.Run("localhost:" + config.Port))
}
