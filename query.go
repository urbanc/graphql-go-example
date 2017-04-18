package main

//import (
//	"strconv"
//
//	"github.com/graphql-go/graphql"
//)
//
//var QueryType = graphql.NewObject(graphql.ObjectConfig{
//	Name: "Query",
//	Fields: graphql.Fields{
//		"allUsers": &graphql.Field{
//			Type:        graphql.NewList(UserType),
//			Description: "List of users",
//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//				list, err := AllUsers()
//				if err != nil {
//					return nil, err
//				}
//
//				return list, nil
//			},
//		},
//
//		"allPosts": &graphql.Field{
//			Type:        graphql.NewList(PostType),
//			Description: "List of posts",
//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//				list, err := AllPosts()
//				if err != nil {
//					return nil, err
//				}
//
//				return list, nil
//			},
//		},
//
//		"user": &graphql.Field{
//			Type: UserType,
//			Args: graphql.FieldConfigArgument{
//				"id": &graphql.ArgumentConfig{
//					Description: "User ID",
//					Type:        graphql.NewNonNull(graphql.ID),
//				},
//			},
//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
//				i := p.Args["id"].(string)
//				id, err := strconv.Atoi(i)
//				if err != nil {
//					return nil, err
//				}
//				return GetUserByID(id)
//			},
//		},
//	},
//})
