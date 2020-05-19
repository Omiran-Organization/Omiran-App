package main

import (
	"Omiran-App/backend/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/graphql", handler.GraphQLService)
	r.Run()
}
