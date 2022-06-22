package handler

import (
	"net/http"
	"zainokta/twirpexample/service1/rpc/serviceone"
	"zainokta/twirpexample/service2/internal"
	"zainokta/twirpexample/service2/internal/usecase/makeshoe"

	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
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
	_ = makeshoe.New(h.container.ShoeRepository)

	h.router.GET("/shoe", func(ctx *gin.Context) {
		client := serviceone.NewServiceOneProtobufClient("https://29eb-140-213-186-43.ngrok.io", &http.Client{
			Transport: &http2.Transport{},
		})
		resp, err := client.MakeHat(ctx, &serviceone.Size{Inches: 12})
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"shoe": resp,
		})
	})
}
