package main

import (
	"Omiran-App/backend/dbutils"
	"Omiran-App/backend/handler"
	"Omiran-App/backend/redis"
	"flag"

	"github.com/gin-gonic/gin"
)

func init() {
	configFlag := flag.String("config", "config.yaml", "backend configuration file")
	dbutils.Open(*flag)
	handler.InitGQLSchema()
	redis.InitCache()
}

func main() {
	r := gin.Default()
	r.POST("/graphql", handler.GraphQLService)
	r.POST("/create", handler.AccountCreationHandler)
	r.POST("/auth", handler.AuthHandler)
	r.POST("/follow", handler.StartFollowingHandler)
	r.POST("/refresh", handler.RefreshSessionHandler)
	r.POST("/signin", handler.SignInHandler)
	r.Run()
}
