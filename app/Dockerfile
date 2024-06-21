FROM golang

USER root

COPY . /app

WORKDIR /app

#RUN apk add build-base

RUN apt-get install gcc

ENV CGO_ENABLED=1

RUN go build .

CMD [ "go", "run", "." ]