package dbutils

import (
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     uint32 `yaml:"port"`
}

type User struct {
	UUID           string
	Username       string
	Email          string
	Password       string
	Description    string
	ProfilePicture string
}

type Follows struct {
	UUID          string
	UserFollowing string
}
