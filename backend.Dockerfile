FROM golang:1.13.0-alpine:3.9
WORKDIR /home/go/src/Omiran-App
COPY backend/backend.sh .
RUN ./backend.sh
COPY backend .
CMD go run main.go
EXPOSE 8080
# Unsure about how I'm going to address creating database users here so that mysql connection actually works
# Probably going to implement a script that generates a YAML file and db user (with CLI args or user input)  

