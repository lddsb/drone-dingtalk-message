FROM golang AS builder
WORKDIR /app
COPY . .
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -a -o drone-dingtalk .

FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY --from=builder /app/drone-dingtalk /bin
COPY --from=builder /app/tpls /app/tpls

ENTRYPOINT ["/bin/drone-dingtalk"]
