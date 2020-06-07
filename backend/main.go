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
	handler.InitCache()

	r.POST("/graphql", handler.GraphQLService)
	r.POST("/create", handler.AccountCreationHandler)
	//r.POST("/signin")
	r.POST("/auth", handler.AuthHandler)
	r.POST("/follow", handler.StartFollowingHandler)
	handler.Examine()
	r.Run()
}
