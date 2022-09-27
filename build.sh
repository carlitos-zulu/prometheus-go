#!/usr/bin/env bash 
set -xe

export GOPRIVATE="github.com/zuluapp"
export SCOPE=develop-read
export APPLICATION=prometheus-go
export APPLICATION_ID=1234

whoami

cp go.mod ~/
cp go.sum ~/
cp -R cmd ~/
cp -R internal ~/

(
    echo "machine github.com"
    echo "    login $GIT_USER"
    echo "    password $GIT_PWD"
    echo "machine api.github.com"
    echo "    login $GIT_USER"
    echo "    password $GIT_PWD"
) > ~/.netrc

chmod 600 ~/.netrc

cat ~/.netrc

ls -lha ~/.netrc

cd

pwd
ls -lha

# install packages and dependencies
go mod tidy

# build command
go build -o bin/application cmd/api/main.go

bin/application