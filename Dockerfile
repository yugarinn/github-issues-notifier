FROM golang:1.20

ENV GOPATH=/go

RUN mkdir /github-issues-notificator
WORKDIR /github-issues-notificator

ADD go.mod ./go.mod
ADD go.sum ./go.sum

ADD . .

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.10.0/wait /usr/local/bin/wait
RUN chmod +x /usr/local/bin/wait
RUN go mod download && go mod verify

EXPOSE 3000

CMD /usr/local/bin/wait && ./setup.sh
