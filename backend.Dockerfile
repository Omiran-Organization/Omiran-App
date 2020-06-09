FROM alpine:3.7
WORKDIR /home/go/src/Omiran-App
COPY backend /home/go/src/Omiran-App 
RUN apt-get update && apt-get upgrade
RUN apt install golang-go && apt install mysql-server 
RUN ./backend.sh
CMD go run main.go
EXPOSE 8080
# Unsure about how I'm going to address creating database users here so that mysql connection actually works
# Probably going to implement a script that generates a YAML file and db user (with CLI args or user input)  

