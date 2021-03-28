FROM golang:1.12-alpine

WORKDIR /go/src/github.com/FixIT-hackathon/meta-transfer-from
COPY . .
RUN apk update \
    && apk --no-cache --update add build-base
RUN CGO_ENABLED=0 GOOS=linux go build -tags netgo -o /usr/local/bin/meta-transfer-from github.com/FixIT-hackathon/meta-transfer-from

###

FROM alpine:3.9

COPY --from=0 /usr/local/bin/meta-transfer-from /usr/local/bin/meta-transfer-from
RUN apk add --no-cache ca-certificates


