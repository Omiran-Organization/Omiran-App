package dbutils

import (
	"errors"
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
	UUID          uuid.UUID `db:"uuid" json:"uuid"`
	UserFollowing uuid.UUID `db:"user_following" json:"user_following"`
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
	DB.Select(&users, "SELECT uuid, username, email, description, profile_picture FROM User")
	return users
}

// SelectAllFollows returns all Follows from the Follows table
func SelectAllFollows() []Follows {
	follows := []Follows{}
	DB.Select(&follows, "SELECT uuid, user_following FROM Follows")
	return follows
}

// Create creates a new User row
func (u *User) Create() error {
	if len(u.Username) > 40 {
		return errors.New("Username too long (can be maximum 40 characters)")
	}

	query, err := DB.Prepare("INSERT INTO User (uuid, username, email, password, description, profile_picture) VALUES  (?, ?, ?, ?, ?, ?)")
	defer query.Close()

	if err != nil {
		log.Printf("Prepare err: %s\n", err)
		return err
	}

	_, err = query.Exec(u.UUID, u.Username, u.Email, u.Password, u.Description, u.ProfilePicture)
	if err != nil {
		return fmt.Errorf("Username '%s' taken", u.Username)
	}

	return nil
}

// Create creates a new Follows row
func (f *Follows) Create() error {
	// REPLACE so it doesn't fail if it already exist. If it already exist we can just return success again. INSERT ... ON DUPLICATE KEY UPDATE could also be used, but it doesn't matter since the keys are the only values.
	query, err := DB.Prepare("REPLACE INTO Follows (uuid, user_following) VALUES (?, ?)")
	defer query.Close()
	if err != nil {
		return errors.New("SQL statement error")
	}
	_, err = query.Exec(f.UUID, f.UserFollowing)
	if err != nil {
		return errors.New("User or followee does not exist")
	}
	return nil
}

// Auth checks to see if a row exists with certain user credentials
func (u *User) Auth() error {
	err := DB.Get(u, "SELECT uuid, username, email, description, profile_picture FROM User WHERE email = ? AND password = ? LIMIT 1", u.Email, u.Password)
	return err
}
