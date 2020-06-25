package dbutils

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"

	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

var (
	// DB is an instance of sqlx.DB
	DB *sqlx.DB
)

// Errors
var (

	// ErrUnauthorized indicates the user is not authorized
	ErrUnauthorized = errors.New("unauthorized")
	// ErrInternalServer indicates an internal server error
	ErrInternalServer = errors.New("internal server error")
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
	Follower uuid.UUID `db:"follower" json:"follower"`
	Followee uuid.UUID `db:"followee" json:"followee"`
}

// Open is a boilerplate function that handles opening of the database (reading credentials from a yaml file as well to open said database)
func Open() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	DB, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(full_db_mysql:%s)/Omiran", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT")))

// 	DB, err = sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(localhost:%d)/Omiran", infostruct.User, infostruct.Password, infostruct.Port))

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

// HashPassword hashes password
func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	return string(bytes), err
}

// Create creates a new User row
func (u *User) Create() error {
	if len(u.Username) > 40 {
		return errors.New("Username too long (can be maximum 40 characters)")
	}

	hashedPassword, err := hashPassword(u.Password)
	u.Password = string(hashedPassword)

	if err != nil {
		log.Printf("Error creating hash: %s\n", err)
		return err
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
	query, err := DB.Prepare("REPLACE INTO Follows (follower, followee) VALUES (?, ?)")
	defer query.Close()
	if err != nil {
		return errors.New("SQL statement error")
	}
	_, err = query.Exec(f.Follower, f.Followee)
	if err != nil {
		return errors.New("User or followee does not exist")
	}
	return nil
}

// Delete deletes a Follows row
func (f *Follows) Delete() error {
	query, err := DB.Prepare("DELETE FROM Follows WHERE follower = ? AND followee = ?")
	defer query.Close()
	if err != nil {
		return errors.New("SQL statement error")
	}
	_, err = query.Exec(f.Follower.String, f.Followee.String)
	if err != nil {
		return errors.New("Unable to delete Follows row")
	}
	return nil
}

// checkPasswordHash checks whether string input hashes to password after extracating salt
func checkPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// Auth returns the user if it exists and password matches
// Returns an empty user in case of error.
func Auth(username string, password string) (User, error) {
	var user User

	err := DB.Get(&user, "SELECT uuid, username, password, email, description, profile_picture FROM User WHERE username = ? LIMIT 1", username)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("user auth err %s\n", err)
		// Problem with query
		return User{}, ErrInternalServer

	} else if err == sql.ErrNoRows {
		// Username does not exist
		return User{}, ErrUnauthorized
	}

	match := checkPasswordHash(password, user.Password)
	if match {
		user.Password = "" // Password not needed outside of this

		return user, nil
	}

	return User{}, ErrUnauthorized
}

// GetFollowers returns a list of all followers of the user passed in.
func GetFollowers(uuid uuid.UUID) ([]User, error) {
	queryString := `
		SELECT User.uuid, username, email FROM User
		JOIN Follows ON User.uuid = Follows.follower 
		WHERE Follows.followee = ?
		`
	var followers []User
	err := DB.Select(&followers, queryString, uuid)
	return followers, err
}

// GetUsersBeingFollowed returns a list of all users you are following.
func GetUsersBeingFollowed(uuid uuid.UUID) ([]User, error) {
	queryString := `
		SELECT User.uuid, username, email FROM User
		JOIN Follows ON User.uuid = Follows.followee 
		WHERE Follows.follower = ?
		`
	var followees []User
	err := DB.Select(&followees, queryString, uuid)
	return followees, err
}

// CreateNewStreamKey creates a new private stream key. If one already
// exists it is overwritten. Then it returns the new
func CreateNewStreamKey(id uuid.UUID) (uuid.UUID, error) {
	streamKey := uuid.NewV4()

	stmnt, err := DB.Prepare("UPDATE User SET private_stream_key = ? WHERE uuid = ?")
	defer stmnt.Close()

	if err != nil {
		return streamKey, ErrInternalServer
	}

	_, err = stmnt.Exec(streamKey, id)

	if err != nil {
		return streamKey, ErrInternalServer
	}

	return streamKey, nil
}

// AuthStreamKey checks if the name and streamkey exists in the database
// This is used to authenticate a stream request.
func AuthStreamKey(name string, privateKey string) error {
	var user User

	query := `
	SELECT uuid FROM User 
	WHERE username = ? 
	AND private_stream_key = ? 
	LIMIT 1
	`

	err := DB.Get(&user, query, name, privateKey)

	if err == sql.ErrNoRows {
		log.Printf("Streamer '%s' with key '%s' not authorized", name, privateKey)
		return ErrUnauthorized

	} else if err != nil {
		return ErrInternalServer
	}

	return nil
}
