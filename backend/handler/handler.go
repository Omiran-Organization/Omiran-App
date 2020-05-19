package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/go"
	"github.com/graphql-go/graphql"
)

func GraphQLService(c *gin.Context) {
	fields := graphql.Fields{}
	rootQuery := graphql.ObjectConfig{}
	schemaConfig := graphql.SchemaConfig{}
	schema, err := graphql.NewSchema()
}
