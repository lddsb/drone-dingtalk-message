FROM golang
WORKDIR /app
COPY . .
RUN GO111MODULE=on go build -o drone-dingtalk-message .

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=0 /app/drone-dingtalk-message /bin
ENTRYPOINT ["/bin/drone-dingtalk-message"]
