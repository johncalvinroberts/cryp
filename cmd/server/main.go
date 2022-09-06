package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

	"github.com/johncalvinroberts/cryp/app/client"
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
	api := router.Group("/api")
	api.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, map[string]interface{}{
			"success": true,
		})
	})
	log.Fatal(router.Run("localhost:" + port))
}
