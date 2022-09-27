#!/usr/bin/env bash 
set -xe

export GOPRIVATE=github.com/zuluapp
export SCOPE=develop-read
export APPLICATION=prometheus-go
export APPLICATION_ID=1234

echo "machine github.com\nlogin $GIT_USER\npassword $GIT_PWD\n\nmachine api.github.com\nlogin $GIT_USER\npassword $GIT_PWD\n" >> ~/.netrc

chmod 600 ~/.netrc

cat ~/.netrc

# install packages and dependencies
go mod tidy

# build command
go build -o bin/application cmd/api/main.go

bin/application