package handler

import (
	"Omiran-App/backend/dbutils"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func GraphQLService(c *gin.Context) {
	if c.Request.Body == nil {
		c.Error("404")
	}
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
