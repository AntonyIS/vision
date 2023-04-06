package http

import (
	"net/http"

	"github.com/AntonyIS/vision1.0/config"
	"github.com/gin-gonic/gin"
)

func RunServer(conf *config.BaseConfig) {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Vision1.0",
		})
	})

	router.Run()
}
