FROM alpine:3.7
WORKDIR /home/go/src/Omiran-App
COPY backend /home/go/src/Omiran-App 
RUN apt-get update && apt-get upgrade
RUN apt install golang-go && apt install mysql-server 
RUN ./backend.sh


