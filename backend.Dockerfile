FROM alpine:3.7 
RUN apt-get update && apt-get upgrade
RUN apt install golang-go && apt install mysql
RUN ./backend.sh
