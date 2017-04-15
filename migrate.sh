#!/bin/bash
set -e
cd "$(dirname "$0")"/
migrate -url postgres://vagrant:vagrant@localhost:5432/graphql?sslmode=disable -path ./migrations $@
