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
	Query          string                 `json:"query"`
	VariableValues map[string]interface{} `json:"variables"`
}

// AccountCreationInput is the data sent when creating an account
type AccountCreationInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// GraphQLService is the handler for GraphQL api
func GraphQLService(c *gin.Context) {
	var q Query
	err := c.BindJSON(&q)
	if err != nil {
		log.Fatalf("Error parsing JSON request body %s", err)
	}
	c.JSON(200, processQuery(q.Query, q.VariableValues))
}

func processQuery(query string, variables map[string]interface{}) *graphql.Result {
	params := graphql.Params{Schema: schema, RequestString: query, VariableValues: variables}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Printf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r
}

// AccountCreationHandler generates a new UUID, receives form values, and creates a new user
func AccountCreationHandler(c *gin.Context) {
	var input AccountCreationInput
	err := c.BindJSON(&input)
	if err != nil {
		log.Fatalf("Error parsing JSON request body %s", err)
	}
	u := uuid.NewV4()
	userIntermediary := &dbutils.User{UUID: u, Username: input.Username, Email: input.Email, Password: input.Password}

	// Maybe 500 status code
	err = userIntermediary.Create()
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.JSON(200, userIntermediary)
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

// RefreshSessionHandler calls refresh cookie from redis and assigns new cookie at /refresh
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

// CreateFollowsHandler handles a request and accordingly creates a follows table row
func CreateFollowsHandler(c *gin.Context) {
	followsStruct := &dbutils.Follows{uuid.FromStringOrNil(c.Request.FormValue("follower")), uuid.FromStringOrNil(c.Request.FormValue("followee"))}
	err := followsStruct.Create()
	if err != nil {
		panic(err)
	}

}

// DeleteFollowsHandler handle a request and accordingly deletes a follows table row
func DeleteFollowsHandler(c *gin.Context) {
	followsStruct := &dbutils.Follows{uuid.FromStringOrNil(c.Request.FormValue("follower")), uuid.FromStringOrNil(c.Request.FormValue("followee"))}
	err := followsStruct.Delete()
	if err != nil {
		panic(err)
	}

}
