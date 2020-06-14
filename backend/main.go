package main

import (
	"Omiran-App/backend/dbutils"
	"Omiran-App/backend/handler"
	"Omiran-App/backend/redis"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	dbutils.Open("config.yaml")
	handler.InitGQLSchema()
	redis.InitCache()
}

func main() {
	os.Setenv("PORT", "9090")
	r := gin.Default()
	r.POST("/graphql", handler.GraphQLService)
	r.POST("/create", handler.AccountCreationHandler)
	r.POST("/auth", handler.AuthHandler)
	r.POST("/follow", handler.StartFollowingHandler)
	r.POST("/refresh", handler.RefreshSessionHandler)
	r.POST("/signin", handler.SignInHandler)
	r.DELETE("/signout", handler.SignOut)
	r.POST("/streamauth", handler.StreamAuth)
	r.Run()
}
