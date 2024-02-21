package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rpolnx/go-background-processor/internal/configs"
)

func InitializeServer(config *configs.AppConfig) (*gin.Engine, error) {
	r := gin.New()

	r.Use(gin.Recovery())

	// r.Use(nil){}

	r.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "OK",
		})
	})

	return r, nil
}
