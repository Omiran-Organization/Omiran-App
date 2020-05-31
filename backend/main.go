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
	r.POST("/graphql", handler.GraphQLService)
	r.POST("/create", handler.AccountCreationHandler)
	r.Run()
}
