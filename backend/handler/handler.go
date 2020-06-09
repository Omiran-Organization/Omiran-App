package handler

import (
	"Omiran-App/backend/dbutils"
	"Omiran-App/backend/gql"

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

// AccountCreationHandler generates a new UUID, receives form values, and creates a new user (auth logic for credentials and stuff will probably happen on the frontend)
func AccountCreationHandler(c *gin.Context) {
	u := uuid.NewV4()
	userIntermediary := &dbutils.User{UUID: u, Username: c.Request.FormValue("username"), Email: c.Request.FormValue("email"), Password: c.Request.FormValue("password"), Description: c.Request.FormValue("description"), ProfilePicture: c.Request.FormValue(("profile_picture"))}

	//Maybe 500 status code
	err := userIntermediary.Create()
	if err != nil {
		c.String(400, err.Error())
		return
	}

	c.String(200, "Success")
}

// StartFollowingHandler handles follow requests
func StartFollowingHandler(c *gin.Context) {
	var follow dbutils.Follows
	err := c.BindJSON(&follow)
	if err != nil {
		c.String(400, "Bad format. Expected {\"follower\": user_uuid, \"followee\": followee_id}")
		return
	}

	err2 := follow.Create()
	if err2 != nil {
		c.String(400, err2.Error())
		return
	}

	c.String(200, "Success")
}

// AuthHandler handles authentication by receiving form values, calling dbutils code, and checking to see if dbutils throws ErrNoRows (if it does, deny access)
func AuthHandler(c *gin.Context) {
	_, err := dbutils.Auth(c.Request.FormValue("username"), c.Request.FormValue("password"))

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
