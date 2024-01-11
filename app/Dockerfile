############################
# STEP 1 build executable binary -> 1.56GB
############################
FROM golang:1.21.0 as builder


WORKDIR /app
COPY go.mod go.sum ./

# curl -v "http://127.0.0.1:8091/health"
RUN apt-get update && apt-get install -y wget ca-certificates gnupg \
&& apt-get -y update \
&& apt-get -y install curl \
&& apt install gettext-base

RUN go mod tidy

COPY . /app

ENV GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64

RUN go build -o bin /app/cmd

############################
# STEP 2 build a small image -> 31.69MB
############################
FROM scratch
WORKDIR /app/cmd
# Copy our static executable.
COPY --from=builder /app /app
#COPY --from=builder /app/bin /app/bin
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

#ENV SERVICE_ENVIRONMENT=development
ENV TEST_KEY=cameFromDockerfile-testKey
ENV service.logLevel=panic
ARG SERVER_PORT=8091
ARG SERVER_ADRR=127.0.0.1
ENV PORT $SERVER_PORT
ENV HOST $SERVER_ADRR
EXPOSE $PORT

# Run the app binary.
CMD ["/app/bin"]