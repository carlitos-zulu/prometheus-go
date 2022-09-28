#!/usr/bin/env bash 
set -xe

export GOPRIVATE="github.com/zuluapp"
export SCOPE=develop-read
export APPLICATION=prometheus-go
export APPLICATION_ID=1234

yum install gh -y

# curl -fsSL https://cli.github.com/packages/githubcli-archive-keyring.gpg | sudo dd of=/usr/share/keyrings/githubcli-archive-keyring.gpg \
# && sudo chmod go+r /usr/share/keyrings/githubcli-archive-keyring.gpg \
# && echo "deb [arch=$(dpkg --print-architecture) signed-by=/usr/share/keyrings/githubcli-archive-keyring.gpg] https://cli.github.com/packages stable main" | sudo tee /etc/apt/sources.list.d/github-cli.list > /dev/null \
# && sudo apt update \
# && sudo apt install gh -y

git config --global user.name $GIT_USER
gh auth login --with-token $GIT_PWD

# whoami

# cp go.mod ~/
# cp go.sum ~/
# cp -R cmd ~/
# cp -R internal ~/

# (
#     echo "machine github.com"
#     echo "    login $GIT_USER"
#     echo "    password $GIT_PWD"
#     echo "machine api.github.com"
#     echo "    login $GIT_USER"
#     echo "    password $GIT_PWD"
# ) > ~/.netrc

# chmod 600 ~/.netrc

# cat ~/.netrc

# ls -lha ~/.netrc

# cd

# pwd
# ls -lha

# install packages and dependencies
go mod tidy

# build command
go build -o bin/application cmd/api/main.go

bin/application