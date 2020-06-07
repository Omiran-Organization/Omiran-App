package gql

import (
	"Omiran-App/backend/dbutils"
	"log"

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

// GraphQLSchema is the schema for the graphql endpoint of Omiran
func GraphQLSchema() graphql.Schema {
	fields := graphql.Fields{
		"Users": &graphql.Field{
			Type:        graphql.NewList(userType),
			Description: "All Users",
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				users := dbutils.SelectAllUsers()
				return users, nil
			},
		},

		"User": &graphql.Field{
			Type:        userType,
			Description: "get users by uuid or username",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"username": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var user dbutils.User
				if name, ok := params.Args["username"]; ok {
					err := dbutils.DB.Get(&user, "SELECT uuid, username, email, description, profile_picture FROM User WHERE username = ?", name)
					if err != nil {
						return nil, nil
					}

				} else if id, ok := params.Args["uuid"].(string); ok {
					uuid, err := uuid.FromString(id)
					if err != nil {
						return nil, nil
					}
					err = dbutils.DB.Get(&user, "SELECT uuid, username, email, description, profile_picture FROM User WHERE uuid = ?", uuid)
					if err != nil {
						return nil, nil
					}
				}
				return user, nil
			},
		},

		"Follows": &graphql.Field{
			Type:        graphql.NewList(followsType),
			Description: "get a list of users you are following or users following you",
			Args: graphql.FieldConfigArgument{
				"uuid": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
				"user_following": &graphql.ArgumentConfig{
					Type: graphql.String,
				},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				// follows := dbutils.SelectAllFollows()
				var followList []dbutils.Follows
				if id, ok := params.Args["uuid"].(string); ok {
					uuid, err := uuid.FromString(id)
					if err != nil {
						return nil, nil
					}
					err = dbutils.DB.Select(&followList, "SELECT uuid, user_following FROM Follows WHERE uuid = ?", uuid)
					if err != nil {
						return nil, nil
					}

				} else if id, ok := params.Args["user_following"].(string); ok {
					uuid, err := uuid.FromString(id)
					if err != nil {
						return nil, nil
					}
					err = dbutils.DB.Select(&followList, "SELECT uuid, user_following FROM Follows WHERE user_following = ?", uuid)
					if err != nil {
						return nil, nil
					}
				}

				return followList, nil
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
