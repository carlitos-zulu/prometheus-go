FROM golang:1.19 AS build

ARG GIT_USER=""
ARG GIT_PWD=""

# git is required to fetch go dependencies
RUN apt-get update && apt-get install -y ca-certificates git-core ssh

# Create the user and group files that will be used in the running 
# container to run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Create a netrc file using the credentials specified using --build-arg
RUN printf "machine github.com\n\
    login ${GIT_USER}\n\
    password ${GIT_PWD}\n\
    \n\
    machine api.github.com\n\
    login ${GIT_USER}\n\
    password ${GIT_PWD}\n"\
    >> /root/.netrc
RUN chmod 600 /root/.netrc

ENV GOPRIVATE="github.com/zuluapp"

# Set the working directory outside $GOPATH to enable the support for modules.
WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/api/main.go

CMD ["app"]

# Run the script: docker build --tag docker-prometheus-go --build-arg GIT_USER=YOUR_MAIL --build-arg GIT_PWD=YOUR_TOKEN .
# After, run: docker run -d -p 8080:8080 docker-prometheus-go
