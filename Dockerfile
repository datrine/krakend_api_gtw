FROM alpinelinux/golang

USER root

COPY . ./app

WORKDIR ./app

RUN apk add build-base

ENV CGO_ENABLED=1

RUN go build cmd/main.go

CMD [ "go", "run", "cmd/main.go" ]