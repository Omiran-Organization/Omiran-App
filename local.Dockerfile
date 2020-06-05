FROM alpine:3.7 

RUN ./backend.sh
RUN ./npm-run.sh