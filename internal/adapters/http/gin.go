package http

import (
	"fmt"
	"net/http"

	"github.com/AntonyIS/vision/config"
	"github.com/gin-gonic/gin"
)

func RunServer(conf *config.BaseConfig) {
	router := gin.Default()

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Vision",
		})
	})
	router.Run(fmt.Sprintf(":%s", conf.Port))
}
