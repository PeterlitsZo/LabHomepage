package main

import (
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	if os.Getenv("RUN_MODE") == "dev" {
		config := cors.DefaultConfig()
		config.AllowAllOrigins = true
		config.AllowHeaders = append(config.AllowHeaders, "Authorization")
		r.Use(cors.New(config))
	}
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	Init()
	register(r)

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
