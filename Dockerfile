FROM golang AS builder
WORKDIR /app
COPY . .
ENV GO111MODULE on
ENV CGO_ENABLED 0
ENV GOOS linux
RUN go build -a -o drone-dingtalk .


FROM alpine:latest
RUN apk update && \
  apk add \
    ca-certificates && \
  rm -rf /var/cache/apk/*

COPY --from=builder /app/drone-dingtalk /bin/
ENTRYPOINT ["/bin/drone-dingtalk"]