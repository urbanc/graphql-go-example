#!/usr/bin/env bash
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {createUser(email:"1@x.co"){id, email}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {createUser(email:"2@y.co"){id, email}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {follow(follower:1, followee:2)}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {unfollow(follower:1, followee:2)}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:2){followers{id, email}}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:1){followers{id, email}}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:2){follower(id:1){ email}}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:1){followees{email}}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:1){followee(id:2){email}}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {createPost(user:1,title:"p1",body:"b1"){id}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {createComment(user:1,post:1,title:"t1",body:"b1"){id}}'
# comment to easy insert example data
#curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {removeComment(id:1)}'
#curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {removePost(id:1)}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:1){post(id:2){title,body}}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:1){posts{id,title,body}}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:1){post(id:2){user{id,email}}}}'

# get list of all users
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{allUsers{id email}}'
