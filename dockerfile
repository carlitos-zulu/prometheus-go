FROM golang:1.19

WORKDIR /usr/src/app

RUN git config --global --add url."git@github.com:".insteadOf "https://github.com/"
ENV GOPRIVATE="github.com/zuluapp"

# install git
RUN apt-get update
RUN apt-get install -y git ssh

# add credentials on build
ARG SSH_PRIVATE_KEY
RUN mkdir /root/.ssh/
RUN echo "${SSH_PRIVATE_KEY}" > /root/.ssh/id_rsa
RUN chmod 600 /root/.ssh/id_rsa

# make sure your domain is accepted
RUN touch /root/.ssh/known_hosts
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app ./cmd/api/main.go

CMD ["app"]
