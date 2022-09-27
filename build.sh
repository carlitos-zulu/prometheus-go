#!/usr/bin/env bash 
set -xe

export GOPRIVATE=github.com/zuluapp
export SCOPE=develop-read
export APPLICATION=prometheus-go
export APPLICATION_ID=1234

apt install gh

git config --global credential.cacheOptions "--timeout 300"
git config --global user.name $GIT_USER
gh auth login --with-token $GIT_PWD

# install packages and dependencies
go mod tidy

# build command
go build -o bin/application cmd/api/main.go

bin/application