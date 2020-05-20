package main

import (
	"Omiran-App/backend/dbutils"
	"Omiran-App/backend/handler"

	"github.com/gin-gonic/gin"
)

func init() {
	dbutils.Open("database-config.yaml")
}

func main() {
	r := gin.Default()
	r.GET("/graphql", handler.GraphQLService)
	r.Run()
}
