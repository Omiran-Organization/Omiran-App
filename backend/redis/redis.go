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
	conn, err := redis.DialURL("redis://redis")
	if err != nil {
		panic(err)
	}
	Cache = conn
}

// SetCachePlusToken sets the cache
func SetCachePlusToken(c *gin.Context, id uuid.UUID) (string, error) {
	sessionToken := uuid.NewV4().String()
	_, err := Cache.Do("SETEX", sessionToken, "1209600", id.String())
	c.SetCookie("session_token", sessionToken, 1209600, "/", "localhost", false, false)

	if err != nil {
		return "", dbutils.ErrInternalServer
	}
	return sessionToken, nil

}

// CheckSessCookie checks cookie when authorizing
func CheckSessCookie(c *gin.Context) error {
	cookie, err := c.Request.Cookie("session_token")
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
	_, err = Cache.Do("SETEX", newSessionToken, "1209600", response)
	if err != nil {
		return dbutils.ErrInternalServer

	}
	_, err = Cache.Do("DEL", sessionToken)
	if err != nil {
		return dbutils.ErrInternalServer
	}
	c.SetCookie("session_token", newSessionToken, 1209600, "/", "localhost", false, false)
	return nil

}

// DeleteSessionByToken deletes a session by token
func DeleteSessionByToken(c *gin.Context) error {
	r := c.Request
	cookie, err := r.Cookie("session_token")
	if err != nil {
		return dbutils.ErrUnauthorized
	}
	sessionToken := cookie.Value
	_, err = Cache.Do("DEL", sessionToken)
	if err != nil {
		return dbutils.ErrInternalServer
	}
	c.SetCookie("session_token", sessionToken, -1, "/", "localhost", false, false)
	return nil
}

// GetLoggedInUUID gets the uuid of the currently logged in user.
// Can be used to check if the user is currently logged in.
func GetLoggedInUUID(c *gin.Context) (uuid.UUID, error) {
	// func GetUser(c *gin.Context) {

	cookie, err := c.Request.Cookie("session_token")
	if err != nil {
		return uuid.UUID{}, dbutils.ErrUnauthorized
	}

	sessionToken := cookie.Value

	idString, err := redis.String(Cache.Do("GET", sessionToken))

	if err != nil {
		log.Println(err)
		return uuid.UUID{}, dbutils.ErrInternalServer
	}

	UUID, err := uuid.FromString(idString)

	if err != nil {
		return uuid.UUID{}, dbutils.ErrUnauthorized
	}

	return UUID, nil
}
