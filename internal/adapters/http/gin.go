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
	usersRoute := router.Group("/v1/users")
	devicesRoute := router.Group("/v1/devices")
	clientsRoute := router.Group("/v1/clients")
	{
		usersRoute.POST("/", handler.PostUser)
		usersRoute.GET("/", handler.GetUsers)
		usersRoute.GET("/:id", handler.GetUser)
		usersRoute.PUT("/:id", handler.PutUser)
		usersRoute.DELETE("/:id", handler.DeleteUser)
	}
	{
		devicesRoute.POST("/", handler.PostController)
		devicesRoute.GET("/", handler.GetControllers)
		devicesRoute.GET("/:id", handler.GetController)
		devicesRoute.PUT("/:id", handler.PutController)
		devicesRoute.DELETE("/:id", handler.DeleteController)
	}
	{
		clientsRoute.POST("/", handler.PostClient)
		clientsRoute.GET("/", handler.GetClients)
		clientsRoute.GET("/:id", handler.GetClient)
		clientsRoute.PUT("/:id", handler.PutClient)
		clientsRoute.DELETE("/:id", handler.DeleteClient)
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Vision",
		})
	})
	router.Run(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
