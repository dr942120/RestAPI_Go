FROM golang:1.19.0

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

# to install and clean up our dependencies
RUN go mod tidy
