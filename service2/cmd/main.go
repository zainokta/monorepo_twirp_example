package main

import (
	"net/http"
	"zainokta/twirpexample/service2/internal"
	"zainokta/twirpexample/service2/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	container := internal.NewContainer()
	router := handler.NewRouter(server, container)
	router.RegisterRouter()

	http.ListenAndServe(":8000", server)
}
