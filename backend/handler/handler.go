package handler

import (
	"Omiran-App/backend/dbutils"
	"encoding/json"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

var userType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"uuid": &graphql.Field{
				Type: graphql.Int,
			},
			"username": &graphql.Field{
				Type: graphql.String,
			},
			"email": &graphql.Field{
				Type: graphql.String,
			},
			"password": &graphql.Field{
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
		"Follows": &graphql.Field{
			Type:        graphql.NewList(followsType),
			Description: "All follows",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return follows, nil
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
