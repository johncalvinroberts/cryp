package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/johncalvinroberts/cryp/internal/client"
	"github.com/johncalvinroberts/cryp/internal/health"
)

const defaultPort = "9000"

func main() {
	fmt.Println("Starting Server")
	router := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	router.Use(static.Serve("/", client.GetClientFileSystem()))
	router.GET("/api/health", GetHealth)
	log.Fatal(router.Run("localhost:" + port))
}

func GetHealth(c *gin.Context) {
	healthy := health.GetHealth()
	c.JSON(http.StatusOK, map[string]interface{}{
		"success": healthy,
	})
}
