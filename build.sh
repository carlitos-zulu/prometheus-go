#!/usr/bin/env bash 
set -xe

export GOPRIVATE=github.com/zuluapp
export SCOPE=develop-read
export APPLICATION=prometheus-go
export APPLICATION_ID=1234

echo "machine github.com\n\
    login $GIT_USER\n\
    password $GIT_PWD\n\
\n\
machine api.github.com\n\
    login $GIT_USER\n\
    password $GIT_PWD\n"\
    >> ~/.netrc

# install packages and dependencies
go mod tidy

# build command
go build -o bin/application cmd/api/main.go

bin/application