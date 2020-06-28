package handler

import (
	"Omiran-App/backend/dbutils"
	"Omiran-App/backend/redis"

	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

// StartStreamAuth authorizes streamer
// This is only intended to be accesses via a stream request sent
// from the nginx stream server. Do not access this from the frontend.
func StartStreamAuth(c *gin.Context) {
	body, _ := ioutil.ReadAll(c.Request.Body)
	fmt.Print("BODY: ")
	fmt.Println(string(body))
	values, err := url.ParseQuery(string(body))

	if err != nil {
		c.String(500, "Internal server error")
		return
	}

	key := values.Get("psk")   // psk = Private Stream Key
	name := values.Get("name") // username associated with stream key

	if key == "" {
		c.String(400, "No key specified")
		return
	}

	err = dbutils.AuthStreamKey(name, key)

	if err != nil {
		c.String(400, "Invalid stream key")
		return
	}

	c.String(200, "Success")
}

// UserStreamKey fetches the streamkey from db
type UserStreamKey struct {
	UUID             uuid.UUID `db:"uuid"`
	Username         string    `db:"username"`
	PrivateStreamKey string    `db:"private_stream_key"`
}

// GetStreamKey returns the streamkey of the current user. This will be the key
// that is checked in StreamAuth. This streamkey that is returned will include
// the stream name. `username?psk=private_stream_key_hash`
func GetStreamKey(c *gin.Context) {

	// If logged in, fetch stream key for that user.

	uuid, err := redis.GetLoggedInUUID(c.Request)

	if err != nil {
		c.String(400, "unauthorized")
		return
	}

	var user UserStreamKey

	err = dbutils.DB.Get(&user, "SELECT uuid, username, private_stream_key FROM User WHERE uuid = ?", uuid.String())

	if err != nil {
		log.Println(err)
		c.String(500, "internal server error")
		return
	}

	if user.PrivateStreamKey == "" {
		newKey, err := dbutils.CreateNewStreamKey(user.UUID)
		if err != nil {
			log.Println(err)
			c.String(500, "internal server error")
			return
		}
		c.String(200, user.Username+"?psk="+newKey.String())
		return
	}

	c.String(200, user.Username+"?psk="+user.PrivateStreamKey)

}

// CreateNewStreamKey creates and returns a new stream key. This can be used to
// create the inital key, or to create a new key in case the old one got leaked.
func CreateNewStreamKey(c *gin.Context) {

	// If logged in, create or update key for user.
	uuid, err := redis.GetLoggedInUUID(c.Request)

	if err != nil {
		c.String(400, "unauthorized")
		return
	}

	var user UserStreamKey

	err = dbutils.DB.Get(&user, "SELECT uuid, username, private_stream_key FROM User WHERE uuid = ?", uuid.String())

	if err != nil {
		log.Println(err)
		c.String(500, "internal server error")
		return
	}

	newKey, err := dbutils.CreateNewStreamKey(uuid)

	if err != nil {
		log.Println(err)
		c.String(500, "internal server error")
		return
	}

	c.String(200, user.Username+"?psk="+newKey.String())

}
