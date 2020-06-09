package redis

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	uuid "github.com/satori/go.uuid"
)

//Cache is where redis cache is setup
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

//SetCache sets the cache
func SetCache(username string) {
	log.Printf(username)
	sessionToken := uuid.NewV4().String()
	_, err := Cache.Do("SETEX", sessionToken, "120", username)

	if err != nil {
		return
	}
	return

}

// SetSessCookie sets the cookie
func SetSessCookie(c *gin.Context) {
	c.SetCookie("session_token", session, 120, "/", "localhost", false, true)
	cookie, err := c.Cookie("session_token")
	log.Printf(cookie)
	if err != nil {
		return
	}
	return

}
