# Build stage
FROM golang:1.25-alpine AS builder
WORKDIR /src

# Install dependencies for building
RUN apk add --no-cache git ca-certificates

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /users-service ./

# Final image
FROM alpine:3.18
RUN apk add --no-cache ca-certificates
WORKDIR /
COPY --from=builder /users-service /users-service

EXPOSE 8080

ENV DATABASE_URL="host=db user=postgres password=postgres dbname=usersdb port=5432 sslmode=disable TimeZone=UTC"

ENTRYPOINT ["/users-service"]
