package redis

import (
	"Omiran-App/backend/dbutils"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

// Cache is where redis cache is setup
var Cache redis.Conn
var session string

// InitCache creates the cache
func InitCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	Cache = conn
}

// SetCachePlusToken sets the cache
func SetCachePlusToken(c *gin.Context, username string) {
	log.Printf(username)
	sessionToken := uuid.NewV4().String()
	_, err := Cache.Do("SETEX", sessionToken, "120", username)
	c.SetCookie("session_token", sessionToken, 120000, "/", "localhost", false, false)
	if err != nil {
		return
	}
	return

}

// CheckSessCookie checks cookie when authorizing
func CheckSessCookie(c *gin.Context) error {
	cookie, err := c.Request.Cookie("session_token")
	if err != nil {
		return dbutils.ErrUnauthorized
	}
	sessionToken := cookie.Value
	response, err := Cache.Do("GET", sessionToken)
	// log.Printf(sessionToken)
	// log.Println(response)
	if err != nil {
		return dbutils.ErrInternalServer
	}
	if response == nil {
		return dbutils.ErrUnauthorized
	}
	return nil
}

// Refresh refreshes a users session_token
func Refresh(c *gin.Context) error {
	r := c.Request
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return dbutils.ErrUnauthorized
	}
	sessionToken := cookie.Value
	response, err := Cache.Do("GET", sessionToken)
	if err != nil {
		return dbutils.ErrInternalServer
	}
	if response == nil {
		return dbutils.ErrUnauthorized
	}
	newSessionToken := uuid.NewV4().String()
	_, err = Cache.Do("SETEX", newSessionToken, "120", response)
	if err != nil {
		return dbutils.ErrInternalServer

	}
	_, err = Cache.Do("DEL", sessionToken)
	if err != nil {
		return dbutils.ErrInternalServer
	}
	c.SetCookie("session_token", newSessionToken, 120000, "/", "localhost", false, false)
	return nil

}
