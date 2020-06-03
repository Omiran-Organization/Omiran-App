package handler

import (
	"Omiran-App/backend/dbutils"
	"database/sql"
	"encoding/json"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
	uuid "github.com/satori/go.uuid"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"profile_picture": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

var followsType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Follows",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.String,
			},
			"user_following": &graphql.Field{
				Type: graphql.String,
			},
		},
	},
)

// GraphQLService is the handler for GraphQL api
func GraphQLService(c *gin.Context) {
	var rBody string
	err := json.NewDecoder(c.Request.Body).Decode(&rBody)
	if err != nil {
		log.Fatalf("Error parsing JSON request body %s", err)
	}
	c.JSON(200, processQuery(rBody))
}

func processQuery(query string) *graphql.Result {
	users := dbutils.SelectAllUsers()
	follows := dbutils.SelectAllFollows()
	params := graphql.Params{Schema: graphQLSchema(users, follows), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r
}

func graphQLSchema(user []dbutils.User, follows []dbutils.Follows) graphql.Schema {
	fields := &graphql.Fields{
		"Users": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "All Users",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return user, nil
			},
		},
		"User": &graphql.Field{
			Type:        userType,
			Description: "get users by any field (except password)",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"username": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"email": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"description": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"profile_picture": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
		},
		"Follows": &graphql.Field{
			Type:        graphql.NewList(followsType),
			Description: "All follows",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return follows, nil
			},
		},
		"Follow": &graphql.Field{
			Type:        followsType,
			Description: "get follows by any field",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"user_following": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)
	if err != nil {
		log.Fatalf("failed to create new schema; %s\n", err)
	}
	return schema
}

// AccountCreationHandler generates a new UUID, receives form values, and creates a new user (auth logic for credentials and stuff will probably happen on the frontend)
func AccountCreationHandler(c *gin.Context) {
	u := uuid.NewV4()
	userIntermediary := &dbutils.User{UUID: u, Username: c.Request.FormValue("username"), Email: c.Request.FormValue("email"), Password: c.Request.FormValue("password"), Description: c.Request.FormValue("description"), ProfilePicture: c.Request.FormValue(("profile_picture"))}
	userIntermediary.Create()
}

// AuthHandler handles authentication by receiving form values, calling dbutils code, and checking to see if dbutils throws ErrNoRows (if it does, deny access)
func AuthHandler(c *gin.Context) {
	userIntermediary := &dbutils.User{Email: c.Request.FormValue("email"), Password: c.Request.FormValue("password")}
	err := userIntermediary.Auth()
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("user auth err %s\n", err)
	} else if err == sql.ErrNoRows {
		c.String(401, "unauthorized")
	} else {
		c.String(200, "example content blah blah blah")
	}
}
