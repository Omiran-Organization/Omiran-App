FROM alpine:3.7

RUN apt-get update && apt-get upgrade
RUN apt install npm
RUN ./npm-run.sh