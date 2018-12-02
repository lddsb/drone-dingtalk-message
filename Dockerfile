FROM alpine:latest

RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

ADD drone-dingtalk-message /bin/
ENTRYPOINT ["/bin/drone-dingtalk-message"]