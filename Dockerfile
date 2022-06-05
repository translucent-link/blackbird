FROM alpine:3
WORKDIR /
COPY blackbird.linux blackbird

CMD ["/blackbird"]

EXPOSE 8080
