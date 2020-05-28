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
	UUID           string `db:"uuid" json:"uuid"`
	Username       string `db:"username" json:"username"`
	Email          string `db:"email" json:"email"`
	Password       string `db:"password" json:"password"`
	Description    string `db:"description" json:"description"`
	ProfilePicture string `db:"profile_picture" json:"profile_picture"`
}

type Follows struct {
	UUID          string `db:"uuid" json:"uuid"`
	UserFollowing string `db:"user_following" json:"user_following"`
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

func SelectAllUsers() []User {
	users := []User{}
	DB.Select(&users, "SELECT * FROM User")
	return users
}

func SelectAllFollows() []Follows {
	follows := []Follows{}
	DB.Select(&follows, "SELECT * FROM Follows")
	return follows
}

func (u *User) Create() {
	query, err := DB.Prepare(fmt.Sprintf("INSERT INTO User (uuid, username, email, password, description, profile_picture) VALUES  (%s, %s, %s, %s, %s, %s)", u.UUID, u.Username, u.Email, u.Password, u.Description, u.ProfilePicture))
	if err != nil {
		log.Fatalf("Prepare err: %s\n", err)
	}
	query.Exec()
}

func (f *Follows) Create() {
	query, err := DB.Prepare(fmt.Sprintf("INSERT INTO User (uuid, username, email, password, description, profile_picture) VALUES (%s, %s)", f.UUID, f.UserFollowing))
	if err != nil {
		log.Fatalf("Prepare err: %s\n", err)
	}
	query.Exec()
}
