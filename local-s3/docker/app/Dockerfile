# syntax = docker/dockerfile:1.3-labs

FROM golang:1.17-alpine as builder

ENV TZ=Asia/Tokyo
WORKDIR /go/src/app/

COPY ./ ./

RUN <<EOF
  apk update
  apk add --no-cache git alpine-sdk tzdata
  go mod download
  go mod tidy
  CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
EOF

CMD ["./main"]

FROM alpine:latest as production

RUN apk --no-cache add ca-certificates

ARG APP_USER_NAME='user'
ARG APP_USER_GROUP=$APP_USER_NAME

RUN <<EOF
    addgroup -S $APP_USER_NAME
    adduser -S $APP_USER_NAME -G $APP_USER_GROUP
    echo $APP_USER_NAME ALL=\(ALL\) NOPASSWD:ALL > /etc/sudoers
    echo $APP_USER_NAME:$APP_USER_GROUP | chpasswd
EOF

COPY --from=builder --chown=$APP_USER_NAME:$APP_USER_GROUP /go/src/app/main /app

USER $APP_USER_NAME

CMD ["/app"]
