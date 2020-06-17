FROM golang:1.13.0
COPY . /go/src/app
WORKDIR /go/src/app
RUN go mod tidy
RUN go mod vendor
CMD go run main.go
ENTRYPOINT [ "/go/src/app" ]
EXPOSE 8080