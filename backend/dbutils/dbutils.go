package dbutils

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var (
	DB *sqlx.DB
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

func Open(filename string) {
	infoStruct := &DBConfig{}
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("database info file error: %s\n", err)
	}
	err = yaml.Unmarshal(file, infoStruct)
	if err != nil {
		log.Fatalf("unmarshalling problem: %s\n", err)
	}
	DB, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%d)/Omiran", infoStruct.User, infoStruct.Password, infoStruct.Port))
	if err != nil {
		log.Fatalf("database connection error: %s\n", err)
	}
}
