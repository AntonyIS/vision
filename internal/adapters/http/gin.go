package http

import (
	"fmt"
	"net/http"
	"os"

	"github.com/AntonyIS/vision/internal/core/services"
	"github.com/gin-gonic/gin"
)

func RunServer(svc services.AppService) {
	router := gin.Default()
	handler := NewHandler(&svc)
	router.Group("/v1")

	router.POST("/users", handler.PostUser)
	router.GET("/users", handler.GetUsers)
	router.GET("/users/:id", handler.GetUser)
	router.PUT("/users/:id", handler.PutUser)
	router.DELETE("/users/:id", handler.DeleteUser)
	router.POST("/devices", handler.PostController)
	router.GET("/devices", handler.GetControllers)
	router.GET("/devices/:id", handler.GetController)
	router.PUT("/devices/:id", handler.PutController)
	router.DELETE("/devices/:id", handler.DeleteController)
	router.POST("/clients", handler.PostClient)
	router.GET("/clients", handler.GetClients)
	router.GET("/clients/:id", handler.GetClient)
	router.PUT("/clients/:id", handler.PutClient)
	router.DELETE("/clients/:id", handler.DeleteClient)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Vision",
		})
	})
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
