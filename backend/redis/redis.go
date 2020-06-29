package redis

import (
	"Omiran-App/backend/dbutils"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
)

// Cache is where redis cache is setup
var Cache redis.Conn
var session string

func getEnv() string {
	return os.Getenv("APP_ENV")
}

// InitCache creates the cache
func InitCache() {
	var err error
	if getEnv() == "prod" {
		err = godotenv.Load(".env_prod")
	} else {
		err = godotenv.Load(".env")
	}

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	conn, err := redis.DialURL(fmt.Sprintf("redis://%s", os.Getenv("REDIS")))
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
func GetLoggedInUUID(r *http.Request) (uuid.UUID, error) {
	cookie, err := r.Cookie("session_token")
	log.Println(cookie)
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
