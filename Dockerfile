FROM golang:1.21

ARG RIN_ENV
ARG GOPATH

ENV GOPATH=/go
ENV RIN_ENV=${RIN_ENV}

RUN mkdir /github-issues-notifier
WORKDIR /github-issues-notifier

ADD go.mod ./go.mod
ADD go.sum ./go.sum
ADD . .

RUN go mod download && go mod verify

ENTRYPOINT ./setup.sh
