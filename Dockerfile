FROM golang:latest AS builder

WORKDIR /go/src/github.com/Shopify/goserver

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build server.go

# use alpine go runtime
FROM golang:1.10-alpine
RUN apk update && apk add --no-cache ca-certificates
# Make port 9090 available to the world outside this container
EXPOSE 9090
COPY --from=builder /go/src/github.com/Shopify/goserver/server /app/server

CMD ["/app/server"]
