package handler

import (
	"Omiran-App/backend/dbutils"
	"database/sql"

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
	params := graphql.Params{Schema: graphQLSchema(), RequestString: query}
	r := graphql.Do(params)
	if len(r.Errors) > 0 {
		log.Printf("failed to execute graphql operation, errors: %+v", r.Errors)
	}
	return r
}

func graphQLSchema() graphql.Schema {
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
				users := dbutils.SelectAllUsers()
				if name, ok := params.Args["username"]; ok {
					for _, u := range users {
						if name == u.Username {
							return u, nil
						}
					}
				} else if id, ok := params.Args["uuid"].(string); ok {
					uuid, err := uuid.FromString(id)
					if err != nil {
						return nil, nil
					}
					for _, u := range users {
						if uuid == u.UUID {
							return u, nil
						}
					}
				}
				return nil, nil
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
				follows := dbutils.SelectAllFollows()
				var followList []dbutils.Follows
				if id, ok := params.Args["uuid"].(string); ok {
					uuid, err := uuid.FromString(id)
					if err != nil {
						return nil, nil
					}
					for _, f := range follows {
						if uuid == f.UUID {
							followList = append(followList, f)
						}
					}
				} else if id, ok := params.Args["user_following"].(string); ok {
					uuid, err := uuid.FromString(id)
					if err != nil {
						return nil, nil
					}
					for _, f := range follows {
						if uuid == f.UserFollowing {
							followList = append(followList, f)
						}
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

// AccountCreationHandler generates a new UUID, receives form values, and creates a new user (auth logic for credentials and stuff will probably happen on the frontend)
func AccountCreationHandler(c *gin.Context) {
	u := uuid.NewV4()
	userIntermediary := &dbutils.User{UUID: u, Username: c.Request.FormValue("username"), Email: c.Request.FormValue("email"), Password: c.Request.FormValue("password"), Description: c.Request.FormValue("description"), ProfilePicture: c.Request.FormValue(("profile_picture"))}

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
		c.String(400, "Bad format. Expected {\"uuid\": user_uuid, \"user_following\": followee_id}")
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
	userIntermediary := &dbutils.User{Email: c.Request.FormValue("email"), Password: c.Request.FormValue("password")}
	err := userIntermediary.Auth()

	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("user auth err %s\n", err)
	} else if err == sql.ErrNoRows {
		c.String(401, "unauthorized")
	} else {
		c.String(200, "Success")
	}
}
