FROM alpine:3.7 

RUN ./backend.sh
RUN ./npm-run.sh
RUN apt-get update && apt-get upgrade
RUN apt install npm