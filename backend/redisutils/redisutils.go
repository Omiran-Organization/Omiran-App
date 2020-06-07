package redisutils

import (
	"github.com/gomodule/redigo/redis"
)

var cache redis.Conn

//InitCache creates the cache
func InitCache() {
	conn, err := redis.DialURL("redis://localhost")
	if err != nil {
		panic(err)
	}
	cache = conn
}
