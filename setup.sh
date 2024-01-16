#!/usr/bin/env sh

if [ "$GISN_ENV" = "development" ]; then
    go install github.com/githubnemo/CompileDaemon@latest
    CompileDaemon -log-prefix=false -build "go build -o bin/github-issues-notificator ./main.go" -command "./bin/github-issues-notificator" -exclude-dir=".git"
else
    go build -o bin/github-issues-notificator ./main.go
    ./bin/github-issues-notificator
fi
