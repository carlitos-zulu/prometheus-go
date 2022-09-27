#!/usr/bin/env bash 
set -xe

export GOPRIVATE="github.com/zuluapp"
export SCOPE=develop-read
export APPLICATION=prometheus-go
export APPLICATION_ID=1234

(
    echo "machine github.com"
    echo "    login $GIT_USER"
    echo "    password $GIT_PWD"
    echo "machine api.github.com"
    echo "    login $GIT_USER"
    echo "    password $GIT_PWD"
 ) > /root/.netrc

chmod 600 /root/.netrc

cat /root/.netrc

pwd
ls -lha

# install packages and dependencies
go mod tidy

# build command
go build -o bin/application cmd/api/main.go

bin/application