package handler

import (
	"Omiran-App/backend/dbutils"
	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql"
)

func GraphQLService(c *gin.Context) {

}

func graphQLSchema(user []dbutils.User, follows []dbutils.Follows) graphql.Schema {

}
