package http

import (
	"github.com/AntonyIS/vision/internal/core/services"
	"github.com/gin-gonic/gin"
)

type GinHandler interface {
	PostUser(ctx *gin.Context)
	GetUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
	PutUser(ctx *gin.Context)
	DeleteUser(ctx *gin.Context)
	PostClient(ctx *gin.Context)
	GetClient(ctx *gin.Context)
	GetClients(ctx *gin.Context)
	PutClient(ctx *gin.Context)
	DeleteClient(ctx *gin.Context)
	PostController(ctx *gin.Context)
	GetController(ctx *gin.Context)
	GetControllers(ctx *gin.Context)
	PutController(ctx *gin.Context)
	DeleteController(ctx *gin.Context)
}

type handler struct {
	svc services.AppService
}

func NewHandler(svc *services.AppService) GinHandler {
	return &handler{
		svc: *svc,
	}
}

func (h *handler) PostUser(ctx *gin.Context) {

}

func (h *handler) GetUser(ctx *gin.Context) {

}

func (h *handler) GetUsers(ctx *gin.Context) {

}

func (h *handler) PutUser(ctx *gin.Context) {

}
func (h *handler) DeleteUser(ctx *gin.Context) {

}

func (h *handler) PostClient(ctx *gin.Context) {

}

func (h *handler) GetClient(ctx *gin.Context) {

}

func (h *handler) GetClients(ctx *gin.Context) {

}

func (h *handler) PutClient(ctx *gin.Context) {

}
func (h *handler) DeleteClient(ctx *gin.Context) {

}

func (h *handler) PostController(ctx *gin.Context) {

}

func (h *handler) GetController(ctx *gin.Context) {

}

func (h *handler) GetControllers(ctx *gin.Context) {

}

func (h *handler) PutController(ctx *gin.Context) {

}
func (h *handler) DeleteController(ctx *gin.Context) {

}
