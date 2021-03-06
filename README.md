# graphql-go-example

Example GraphQL API implemented in Go and backed by Postgresql

## How to run it

To run this project you need to:
- install golang, see [this guide](https://golang.org/doc/install)
- install [Masterminds/glide](https://github.com/Masterminds/glide) which is a package manager for golang projects.
- install all the dependencies for this project: `glide install`, glide will store all dependencies in `/vendor` folder.
- install postgresql with [docker](https://www.docker.com/community-edition#/download) (run `docker-compose up -d`). For this application, I created a postgresql user called `vagrant` with `vagrant` as password and the database called `graphql`, but of course you can change these settings in `./migrate.sh` and in `./main.go` files.
- install [mattes/migrate](https://github.com/mattes/migrate) which is a tool to create and run migrations against sql databases.
- run the migrations which will create the database tables and indexes `./migrate.sh up`. If you ever want to clean up the the database run `./migrate.sh down` then `./migrate.sh up` again.
- run GraphQL server with command `go run *.go`

## GraphiQL - [https://github.com/graphql/graphiql](https://github.com/graphql/graphiql)

An in-browser IDE for exploring GraphQL access via GraphQL server `http://localhost:8080/` 

Chome extension UNOFFICIAL! - [https://chrome.google.com/webstore/detail/chromeiql/fkkiamalmpiidkljmicmjfbieiclmeij](https://chrome.google.com/webstore/detail/chromeiql/fkkiamalmpiidkljmicmjfbieiclmeij)

## Commands

This application exposes a single endpoints `/graphql` which accepts both mutations and queries.
The following are examples of curl calls to this endpoint:

```bash
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
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {removeComment(id:1)}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d 'mutation {removePost(id:1)}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:1){post(id:2){title,body}}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:1){posts{id,title,body}}}'
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{user(id:1){post(id:2){user{id,email}}}}'

# get list of all users
curl -XPOST http://localhost:8080/graphql -H 'Content-Type:application/graphql' -d '{allUsers{id email}}'

```

## TODO
- [ ] move server into Docker
- [x] move PostgreSQL into Docker
- [ ] add [Relay](https://facebook.github.io/relay/) support
- [ ] check duplicated resolve logic
- [ ] add Makefile
