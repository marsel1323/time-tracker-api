FROM golang:alpine as builder

# Add Maintainer info
LABEL maintainer="Marsel Abdrakhmanov <amn132313@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed
RUN go mod download

COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/migrations /migrations
#COPY --from=builder /app/.env .

# Expose port 8080 to the outside world
EXPOSE 80

#Command to run the executable
CMD ["./main", "-port=80", "-env=production", "-db-dsn=postgres://user:password@db/timetracker?sslmode=disable"]