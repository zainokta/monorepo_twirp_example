package main

import (
	"net/http"
	"zainokta/twirpexample/service1/internal"
	"zainokta/twirpexample/service1/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	container := internal.NewContainer()
	router := handler.NewRouter(server, container)
	router.RegisterRouter()

	http.ListenAndServe(":8080", server)
}
