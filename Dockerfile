FROM alpine:latest

ADD ./web /web

ENTRYPOINT ["./web"]
