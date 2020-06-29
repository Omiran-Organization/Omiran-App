
# Start from golang base image
FROM golang:alpine as builder

# ENV GO111MODULE=on

# Add Maintainer info
# LABEL maintainer="Steven Victor <chikodi543@gmail.com>"

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Set the current working directory inside the container 
WORKDIR /app
ENV APP_ENV=prod
# Copy go mod and sum files 
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

# Start a new stage from scratch
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /app/main .
COPY --from=builder /app/.env . 
COPY --from=builder /app/.env_prod . 
# COPY --from=builder /app/config.yaml .
COPY --from=builder /app/index.html .
# Expose port 8080 to the outside world
EXPOSE 8080
EXPOSE 5000

#Command to run the executable
CMD ["./main"]



#ALTERNATIVE DOCKERFILE
# FROM golang:alpine AS builder

# # Set necessary environmet variables needed for our image
# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

# # Move to working directory /build
# WORKDIR /build

# # Copy and download dependency using go mod
# COPY go.mod .
# COPY go.sum .
# RUN go mod download

# # Copy the code into the container
# COPY . .

# # Build the application
# RUN go build -o main .

# # Move to /dist directory as the place for resulting binary folder
# WORKDIR /dist

# # Copy binary from build to main folder
# RUN cp /build/main .

# # Build a small image
# FROM scratch

# COPY --from=builder /dist/main /

# # Command to run
# ENTRYPOINT ["/main"]

