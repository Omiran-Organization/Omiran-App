package handler

import (
	"Omiran-App/backend/dbutils"
	"Omiran-App/backend/redis"
	"log"

	"github.com/gin-gonic/gin"
)

// SignInHandler signs in user
func SignInHandler(c *gin.Context) {
	var creds SignInCredentials
	err := c.BindJSON(&creds)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(creds)
	username := creds.Username
	password := creds.Password
	user, err := dbutils.Auth(username, password)
	switch err {
	case dbutils.ErrUnauthorized:
		c.String(401, err.Error())
	case dbutils.ErrInternalServer:
		c.String(500, err.Error())
	case nil:
		token, err := redis.SetCachePlusToken(c, user.UUID)

		log.Println(token)
		if err != nil {
			c.String(500, "Cookie not present")
		}

		var re SignInData
		re.UUID = user.UUID
		re.Email = user.Email
		re.Description = user.Description
		re.Username = user.Username
		re.ProfilePicture = user.ProfilePicture
		c.JSON(200, re)

	default:
		c.String(500, "internal server error")
	}
}
