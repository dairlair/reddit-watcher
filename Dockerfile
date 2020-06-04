FROM alpine:3.12

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates

ENV PROBES_PORT 80

EXPOSE $PROBES_PORT

COPY reddit-watcher /reddit-watcher
CMD ["/reddit-watcher"]
