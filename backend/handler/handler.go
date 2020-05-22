package handler

import (
	"Omiran-App/backend/dbutils"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func GraphQLService(c *gin.Context) {

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
}
