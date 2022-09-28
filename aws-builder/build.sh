#!/usr/bin/env bash 
set -xe

(
    echo "machine github.com"
    echo "    login $GIT_USER"
    echo "    password $GIT_PWD"
    echo "machine api.github.com"
    echo "    login $GIT_USER"
    echo "    password $GIT_PWD"
) > ~/.netrc

chmod 600 ~/.netrc

# install packages and dependencies
go mod tidy

# build command
go build -o bin/application cmd/api/main.go

bin/application
