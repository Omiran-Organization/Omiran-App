FROM golang:1.13.0-alpine3.10
COPY ../backend /go/src/app
RUN go mod tidy
RUN go mod vendor
CMD go run main.go
EXPOSE 8080