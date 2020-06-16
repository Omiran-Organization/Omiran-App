package handler

import (
	"Omiran-App/backend/dbutils"
	"Omiran-App/backend/gql"
	"Omiran-App/backend/redis"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
)

var (
	schema graphql.Schema
)

// InitGQLSchema initializes the schema for graphql.
func InitGQLSchema() {
	schema = gql.GraphQLSchema()
}

// Query is for deserializing graphql queries
type Query struct {
	Query string `json:"query"`
}

//Credentials is for structuring the signin route
type Credentials struct {
	Username string `json:"username" `
	Password string `json:"password"`
}

//SignInData is structured data that will be converted to json and sent bck to the client
type SignInData struct {
	UUID           uuid.UUID `json:"uuid"`
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	Description    string    `json:"description`
	ProfilePicture string    `json:"profilepicture`
	Token          string    `json:"token" `
}

// GraphQLService is the handler for GraphQL api
func GraphQLService(c *gin.Context) {
	var q Query
	err := c.BindJSON(&q)
	if err != nil {
		log.Fatalf("Error parsing JSON request body %s", err)
	}
	c.JSON(200, processQuery(q.Query))
}

func processQuery(query string) *graphql.Result {
	params := graphql.Params{Schema: schema, RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Printf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r
}

// AccountCreationHandler generates a new UUID, receives form values, and creates a new user
func AccountCreationHandler(c *gin.Context) {
	u := uuid.NewV4()
	userIntermediary := &dbutils.User{UUID: u, Username: c.Request.FormValue("username"), Email: c.Request.FormValue("email"), Password: c.Request.FormValue("password"), Description: c.Request.FormValue("description"), ProfilePicture: c.Request.FormValue(("profile_picture"))}

	// Maybe 500 status code
	err := userIntermediary.Create()
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(200, "Success")
}

// func setCookieResponse(c *gin.Context) string {
// 	r := c.Request

// 	cookie, err := r.Cookie("session_token")

// 	// token := cookie.Value
// 	if err != nil {

// 		log.Println(err)
// 	}
// 	log.Println(cookie)
// 	return "hello"
// }

// SignInHandler signs in user
func SignInHandler(c *gin.Context) {

	var creds Credentials
	err := c.BindJSON(&creds)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(creds)
	username := creds.Username
	password := creds.Password
	user, err2 := dbutils.Auth(username, password)
	switch err2 {
	case dbutils.ErrUnauthorized:
		c.String(401, err2.Error())
	case dbutils.ErrInternalServer:
		c.String(500, err2.Error())
	case nil:
		redis.SetCachePlusToken(c, user.UUID)

		// r := c.Request

		// cookie, err := r.Cookie("session_token")

		// token := cookie.Value
		// if err != nil {
		// 	c.String(500, "Cookie not present")
		// }
		// log.Println(cookie)
		var re SignInData
		re.UUID = user.UUID
		re.Email = user.Email
		re.Description = user.Description
		re.Username = user.Username
		re.ProfilePicture = user.ProfilePicture
		// re.Token = token
		c.JSON(200, re)

	default:
		c.String(500, "internal server error")
	}

}

// StartFollowingHandler handles follow requests
func StartFollowingHandler(c *gin.Context) {

	UUID, err := redis.GetLoggedInUUID(c)

	if err != nil {
		c.String(400, "unauthorized")
		return
	}

	var follow dbutils.Follows
	err = c.BindJSON(&follow)
	if err != nil {
		c.String(400, "Bad format. Expected {\"followee\": followee_id}")
		return
	}

	follow.Follower = UUID
	err2 := follow.Create()
	if err2 != nil {
		c.String(400, err2.Error())
		return
	}

	c.String(200, "Success")
}

// AuthHandler handles authentication by receiving form values, calling dbutils code, and checking to see if dbutils throws ErrNoRows (if it does, deny access)
func AuthHandler(c *gin.Context) {

	err := redis.CheckSessCookie(c)
	switch err {
	case dbutils.ErrUnauthorized:
		c.String(401, err.Error())
	case dbutils.ErrInternalServer:
		c.String(500, err.Error())
	case nil:
		c.String(200, "success")
	default:
		c.String(500, "internal server error")
	}

}

//RefreshSessionHandler calls refresh cookie from redis and assigns new cookie at /refresh
func RefreshSessionHandler(c *gin.Context) {
	err := redis.Refresh(c)
	switch err {
	case dbutils.ErrUnauthorized:
		c.String(401, err.Error())
	case dbutils.ErrInternalServer:
		c.String(500, err.Error())
	case nil:
		c.String(200, "success")
	default:
		c.String(500, "internal server error")
	}
}

// SignOut signs a user out by clearing their cookies and deleting their session cache
func SignOut(c *gin.Context) {
	err := redis.DeleteSessionByToken(c)
	switch err {
	case dbutils.ErrUnauthorized:
		c.String(401, err.Error())
	case dbutils.ErrInternalServer:
		c.String(500, err.Error())
	case nil:
		c.String(200, "success")
	default:
		c.String(500, "internal server error")
	}
}
