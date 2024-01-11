FROM golang:1.21.0 as builder


WORKDIR /app
COPY go.mod go.sum ./

# curl -v "http://127.0.0.1:8091/health"
RUN apt-get update && apt-get install -y wget ca-certificates gnupg \
&& apt-get -y update \
&& apt-get -y install curl \
&& apt install gettext-base


ENV GO111MODULE=on

RUN go mod download
RUN go mod verify

COPY . /app

ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

#ENV SERVICE_ENVIRONMENT=development
ENV TEST_KEY=cameFromDockerfile-testKey

ARG SERVER_PORT=8091
ARG SERVER_ADRR=127.0.0.1
ENV PORT $SERVER_PORT
ENV HOST $SERVER_ADRR
EXPOSE $PORT

RUN export TEST_KEY

WORKDIR /app/cmd
#RUN go build -o bin/server main.go
#CMD ./bin/server
RUN go build -o go-rest-api .


ENTRYPOINT ["./go-rest-api"]