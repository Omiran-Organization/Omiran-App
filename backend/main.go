package main

import (
	"Omiran-App/backend/dbutils"
	"Omiran-App/backend/handler"
	"Omiran-App/backend/redis"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	dbutils.Open("config.yaml")
	handler.InitGQLSchema()
	redis.InitCache()
}

func main() {
	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"POST", "GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	// r.Use(cors.Default())
	r.POST("/graphql", handler.GraphQLService)
	r.POST("/create", handler.AccountCreationHandler)
	r.POST("/auth", handler.AuthHandler)
	r.POST("/follow", handler.StartFollowingHandler)
	r.POST("/refresh", handler.RefreshSessionHandler)
	r.POST("/signin", handler.SignInHandler)
	r.DELETE("/signout", handler.SignOut)
	r.Run()
}
