#!/usr/bin/env bash 
set -xe

export GOPRIVATE=github.com/zuluapp
export SCOPE=develop-read
export APPLICATION=prometheus-go
export APPLICATION_ID=1234

# install packages and dependencies
go mod tidy

# build command
go build -o bin/application cmd/api/main.go

bin/application