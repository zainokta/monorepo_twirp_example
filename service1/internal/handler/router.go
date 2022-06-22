package handler

import (
	"net/http"
	"zainokta/twirpexample/service1/internal"
	"zainokta/twirpexample/service1/internal/usecase/makehat"
	"zainokta/twirpexample/service1/rpc/serviceone"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router    *gin.Engine
	container *internal.Container
}

func NewRouter(engine *gin.Engine, container *internal.Container) *Router {
	return &Router{
		router:    engine,
		container: container,
	}
}

func (h *Router) RegisterRouter() {
	makeHatImpl := makehat.New(h.container.HatRepository)
	twirpHandler := serviceone.NewServiceOneServer(makeHatImpl)

	h.router.GET("/", func(ctx *gin.Context) {
		hat := h.container.HatRepository.Make(12)
		ctx.JSON(http.StatusOK, hat)
	})
	h.router.POST("/*any", gin.WrapF(twirpHandler.ServeHTTP))
}
