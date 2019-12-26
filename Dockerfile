# stage 1
FROM golang:1.13-alpine as stage

WORKDIR /gorilla-restsvc
COPY go.mod go.sum ./
RUN go mod download
# copy the source from the current directory to the Working Directory inside the container
COPY . .

ENV GO111MODULE=auto
RUN CGO_ENABLED=0 GOOS=linux go build github.com/isgo-golgo13/gorilla-restsvc/cmd/service

# stage 2
FROM alpine:latest
RUN apk add --no-cache tzdata curl 
COPY --from=stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
WORKDIR /root/
COPY --from=stage /gorilla-restsvc .
# healthcheck
HEALTHCHECK CMD curl --fail http://localhost:8080/ || exit 1
CMD ["./service"]