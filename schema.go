package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/base-dev/graphql"
	"github.com/base-dev/relay"
	"golang.org/x/net/context"
)

var nodeDefinitions *relay.NodeDefinitions
var userType *graphql.Object
var postType *graphql.Object

// exported schema, defined in init()
var Schema graphql.Schema

func init() {

	/**
	 * We get the node interface and field from the relay library.
	 *
	 * The first method is the way we resolve an ID to its object. The second is the
	 * way we resolve an object that implements node to its type.
	 */
	nodeDefinitions = relay.NewNodeDefinitions(relay.NodeDefinitionsConfig{
		IDFetcher: func(id string, info graphql.ResolveInfo, ctx context.Context) (interface{}, error) {
			log.Println("IDFetcher called ...")
			// resolve id from global id
			resolvedID := relay.FromGlobalID(id)

			// based on id and its type, return the object
			switch resolvedID.Type {
			case "User":
				id, err := strconv.Atoi(resolvedID.ID)
				if err != nil {
					log.Fatal(err)
				}

				user, err := GetUserByID(id)
				if err != nil {
					log.Fatal(err)
				}
				log.Print("IDFetcher user data: %#v", user)
				return user, err
			case "Post":
				id, err := strconv.Atoi(resolvedID.ID)
				if err != nil {
					log.Fatal(err)
				}

				post, err := GetPostByID(id)
				if err != nil {
					log.Fatal(err)
				}
				log.Print("IDFetcher post data: %#v", post)
				return post, err
			default:
				return nil, errors.New("Unknown node type")
			}
		},
		TypeResolve: func(p graphql.ResolveTypeParams) *graphql.Object {
			log.Printf("TypeResolve: %#v", p.Value)

			// based on the type of the value, return GraphQLObjectType
			switch p.Value.(type) {
			case *User:
				log.Printf("TypeResolve user: %#v", p.Value)
				return userType
			case *Post:
				log.Printf("TypeResolve post: %#v", p.Value)
				return postType
			default:
				log.Printf("TypeResolve default: %#v", p.Value)
				return userType
			}
		},
	})

	/**
	 * We define our basic ship type.
	 *
	 * This implements the following type system shorthand:
	 *   type Ship : Node {
	 *     id: String!
	 *     name: String
	 *   }
	 */
	//shipType = graphql.NewObject(graphql.ObjectConfig{
	//	Name:        "Ship",
	//	Description: "A ship in the Star Wars saga",
	//	Fields: graphql.Fields{
	//		"id": relay.GlobalIDField("Ship", nil),
	//		"name": &graphql.Field{
	//			Type:        graphql.String,
	//			Description: "The name of the ship.",
	//		},
	//	},
	//	Interfaces: []*graphql.Interface{
	//		nodeDefinitions.NodeInterface,
	//	},
	//})

	//postConnectionDefinition := relay.ConnectionDefinitions(relay.ConnectionConfig{
	//	Name:     "Post",
	//	NodeType: postType,
	//})

	userType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "User",
		Description: "User description here ....",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("User", nil),
			"user_id": &graphql.Field{
				Type:        graphql.String,
				Description: "User ID description here ...",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok == true {
						log.Println("User.ID is NOT nil.")
						return user.ID, nil
					}

					log.Println("User.ID is nil!")
					return nil, nil
				},
			},
			"email": &graphql.Field{
				Type:        graphql.String,
				Description: "Email description here ...",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if user, ok := p.Source.(*User); ok == true {
						log.Println("User.email is NOT nil.")
						return user.Email, nil
					}

					log.Println("User.email is nil!")
					return nil, nil
				},
			},
			//"posts": &graphql.Field{
			//	Type: graphql.NewList(postType),
			//Args: relay.ConnectionArgs,
			//Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			//	// convert args map[string]interface into ConnectionArguments
			//	args := relay.NewConnectionArguments(p.Args)
			//
			//	// get ship objects from current faction
			//	posts := []interface{}{}
			//	// let relay library figure out the result, given
			//	// - and the filter arguments (i.e. first, last, after, before)
			//	return relay.ConnectionFromArray(posts, args), nil
			//},
			//},
		},
		Interfaces: []*graphql.Interface{
			nodeDefinitions.NodeInterface,
		},
	})

	postType = graphql.NewObject(graphql.ObjectConfig{
		Name:        "Post",
		Description: "Post description here ....",
		Fields: graphql.Fields{
			"id": relay.GlobalIDField("Post", nil),
			"post_id": &graphql.Field{
				Type:        graphql.String,
				Description: "Post ID description here ...",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if post, ok := p.Source.(*Post); ok == true {
						log.Println(" Post.ID is NOT nil.")
						return post.ID, nil
					}

					log.Println("Post.ID is nil!")
					return nil, nil
				},
			},
			"title": &graphql.Field{
				Type:        graphql.String,
				Description: "Title description here ...",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if post, ok := p.Source.(*Post); ok == true {
						log.Println("Post.title is NOT nil.")
						return post.Title, nil
					}

					log.Println("Post.title is nil!")
					return nil, nil
				},
			},
			"body": &graphql.Field{
				Type:        graphql.String,
				Description: "Body description here ...",
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					if post, ok := p.Source.(*Post); ok == true {
						log.Println("Post.Body is NOT nil.")
						return post.Body, nil
					}

					log.Println("Post.Body is nil!")
					return nil, nil
				},
			},
		},
	})

	/**
	 * We define a connection between a faction and its ships.
	 *
	 * connectionType implements the following type system shorthand:
	 *   type ShipConnection {
	 *     edges: [ShipEdge]
	 *     pageInfo: PageInfo!
	 *   }
	 *
	 * connectionType has an edges field - a list of edgeTypes that implement the
	 * following type system shorthand:
	 *   type ShipEdge {
	 *     cursor: String!
	 *     node: Ship
	 *   }
	 */

	userConnectionDefinition := relay.ConnectionDefinitions(relay.ConnectionConfig{
		Name:     "User",
		NodeType: userType,
	})

	/**
	 * We define our faction type, which implements the node interface.
	 *
	 * This implements the following type system shorthand:
	 *   type Faction : Node {
	 *     id: String!
	 *     name: String
	 *     ships: ShipConnection
	 *   }
	 */
	//factionType = graphql.NewObject(graphql.ObjectConfig{
	//	Name:        "Faction",
	//	Description: "A faction in the Star Wars saga",
	//	Fields: graphql.Fields{
	//		"id": relay.GlobalIDField("Faction", nil),
	//		"name": &graphql.Field{
	//			Type:        graphql.String,
	//			Description: "The name of the faction.",
	//		},
	//		"ships": &graphql.Field{
	//			Type: shipConnectionDefinition.ConnectionType,
	//			Args: relay.ConnectionArgs,
	//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	//				// convert args map[string]interface into ConnectionArguments
	//				args := relay.NewConnectionArguments(p.Args)
	//
	//				// get ship objects from current faction
	//				ships := []interface{}{}
	//				if faction, ok := p.Source.(*Faction); ok {
	//					for _, shipId := range faction.Ships {
	//						ships = append(ships, GetShip(shipId))
	//					}
	//				}
	//				// let relay library figure out the result, given
	//				// - the list of ships for this faction
	//				// - and the filter arguments (i.e. first, last, after, before)
	//				return relay.ConnectionFromArray(ships, args), nil
	//			},
	//		},
	//	},
	//	Interfaces: []*graphql.Interface{
	//		nodeDefinitions.NodeInterface,
	//	},
	//})

	/**
	 * This is the type that will be the root of our query, and the
	 * entry point into our schema.
	 *
	 * This implements the following type system shorthand:
	 *   type Query {
	 *     rebels: Faction
	 *     empire: Faction
	 *     node(id: String!): Node
	 *   }
	 */
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"users": &graphql.Field{
				Type: userConnectionDefinition.ConnectionType,
				Args: relay.ConnectionArgs,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// convert args map[string]interface into ConnectionArguments
					args := relay.NewConnectionArguments(p.Args)
					log.Printf("args: %#v", args)
					pagination := Arguments{
						After: fmt.Sprintf("%s", args.After),
					}

					users := []interface{}{}
					usersList, err := AllUsers(pagination)
					if err != nil {
						log.Fatal(err)
					}

					for _, user := range usersList {
						users = append(users, user)
					}

					// let relay library figure out the result, given
					// - and the filter arguments (i.e. first, last, after, before)
					return relay.ConnectionFromArray(users, args), nil
				},
			},
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Description: "User ID",
						Type:        graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					i := p.Args["id"].(string)
					id, err := strconv.Atoi(i)
					if err != nil {
						return nil, err
					}
					user, err := GetUserByID(id)
					if err != nil {
						log.Fatal(err)
					}

					return user, nil
				},
			},
			"node": nodeDefinitions.NodeField,
		},
	})

	/**
	 * This will return a GraphQLField for our ship
	 * mutation.
	 *
	 * It creates these two types implicitly:
	 *   input IntroduceShipInput {
	 *     clientMutationID: string!
	 *     shipName: string!
	 *     factionId: ID!
	 *   }
	 *
	 *   input IntroduceShipPayload {
	 *     clientMutationID: string!
	 *     ship: Ship
	 *     faction: Faction
	 *   }
	 */
	//shipMutation := relay.MutationWithClientMutationID(relay.MutationConfig{
	//	Name: "IntroduceShip",
	//	InputFields: graphql.InputObjectConfigFieldMap{
	//		"shipName": &graphql.InputObjectFieldConfig{
	//			Type: graphql.NewNonNull(graphql.String),
	//		},
	//		"factionId": &graphql.InputObjectFieldConfig{
	//			Type: graphql.NewNonNull(graphql.ID),
	//		},
	//	},
	//	OutputFields: graphql.Fields{
	//		"ship": &graphql.Field{
	//			Type: shipType,
	//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	//				if payload, ok := p.Source.(map[string]interface{}); ok {
	//					return GetShip(payload["shipId"].(string)), nil
	//				}
	//				return nil, nil
	//			},
	//		},
	//		"faction": &graphql.Field{
	//			Type: factionType,
	//			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
	//				if payload, ok := p.Source.(map[string]interface{}); ok {
	//					return GetFaction(payload["factionId"].(string)), nil
	//				}
	//				return nil, nil
	//			},
	//		},
	//	},
	//	MutateAndGetPayload: func(inputMap map[string]interface{}, info graphql.ResolveInfo, ctx context.Context) (map[string]interface{}, error) {
	//		// `inputMap` is a map with keys/fields as specified in `InputFields`
	//		// Note, that these fields were specified as non-nullables, so we can assume that it exists.
	//		shipName := inputMap["shipName"].(string)
	//		factionId := inputMap["factionId"].(string)
	//
	//		// This mutation involves us creating (introducing) a new ship
	//		newShip := CreateShip(shipName, factionId)
	//		// return payload
	//		return map[string]interface{}{
	//			"shipId":    newShip.ID,
	//			"factionId": factionId,
	//		}, nil
	//	},
	//})

	/**
	 * This is the type that will be the root of our mutations, and the
	 * entry point into performing writes in our schema.
	 *
	 * This implements the following type system shorthand:
	 *   type Mutation {
	 *     introduceShip(input IntroduceShipInput!): IntroduceShipPayload
	 *   }
	 */

	//mutationType := graphql.NewObject(graphql.ObjectConfig{
	//	Name: "Mutation",
	//	Fields: graphql.Fields{
	//		"introduceShip": shipMutation,
	//	},
	//})

	/**
	 * Finally, we construct our schema (whose starting query type is the query
	 * type we defined above) and export it.
	 */
	var err error
	Schema, err = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
		//Mutation: mutationType,
	})
	if err != nil {
		// panic if there is an error in schema
		log.Fatal(err)
	}
}
