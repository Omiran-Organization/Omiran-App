package handler

import (
	"Omiran-App/backend/dbutils"
	"encoding/json"
	"fmt"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func GraphQLService(c *gin.Context) {

}

func processQuery(query string) (result string) {
	users := dbutils.SelectAllUsers()
	follows := dbutils.SelectAllFollows()
	params := graphql.Params{Schema: graphQLSchema(users, follows), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Fatalf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	rJSON, _ := json.Marshal(r)

	return fmt.Sprintf("%s", rJSON)
}

func graphQLSchema(user []dbutils.User, follows []dbutils.Follows) graphql.Schema {
	fields := &graphql.Fields{
		"Users": &graphql.Field{
			Type:        graphql.NewList(dbutils.User),
			Description: "All Users",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				return user, nil
			},
		},
		"Follows": &graphql.Field{
			Type:        graphql.NewList(dbutils.Follows),
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
