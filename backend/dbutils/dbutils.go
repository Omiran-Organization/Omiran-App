package dbutils

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

var (
	// DB is an instance of sqlx.DB
	DB *sqlx.DB
)

// DBConfig is a database configuration abstraction struct
type DBConfig struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Port     uint32 `yaml:"port"`
}

// User is a sqlx database User table abstraction struct
type User struct {
	UUID           uuid.UUID `db:"uuid" json:"uuid"`
	Username       string    `db:"username" json:"username"`
	Email          string    `db:"email" json:"email"`
	Password       string    `db:"password" json:"password"`
	Description    string    `db:"description" json:"description"`
	ProfilePicture string    `db:"profile_picture" json:"profile_picture"`
}

// Follows is a sqlx database Follows table abstraction struct
type Follows struct {
	UUID          string `db:"uuid" json:"uuid"`
	UserFollowing string `db:"user_following" json:"user_following"`
}

// Open is a boilerplate function that handles opening of the database (reading credentials from a yaml file as well to open said database)
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

// SelectAllUsers returns all Users from the User table
func SelectAllUsers() []User {
	users := []User{}
	DB.Select(&users, "SELECT * FROM User")
	return users
}

// SelectAllFollows returns all Follows from the Follows table
func SelectAllFollows() []Follows {
	follows := []Follows{}
	DB.Select(&follows, "SELECT * FROM Follows")
	return follows
}

// Create creates a new User row
func (u *User) Create() {
	query, err := DB.Prepare(fmt.Sprintf("INSERT INTO User (uuid, username, email, password, description, profile_picture) VALUES  ('%s', '%s', '%s', '%s', '%s', '%s')", u.UUID, u.Username, u.Email, u.Password, u.Description, u.ProfilePicture))
	if err != nil {
		log.Fatalf("Prepare err: %s\n", err)
	}
	query.Exec()
}

// Create creates a new Follows row
func (f *Follows) Create() {
	query, err := DB.Prepare(fmt.Sprintf("INSERT INTO Follows (uuid, user_following) VALUES ('%s', '%s')", f.UUID, f.UserFollowing))
	if err != nil {
		log.Fatalf("Prepare err: %s\n", err)
	}
	query.Exec()
}

// Auth checks to see if a row exists with certain user credentials
func (u *User) Auth() error {
	err := DB.Select(&u, fmt.Sprintf("SELECT * FROM User WHERE email=%s AND WHERE password=%s", u.Email, u.Password))
	return err
}
