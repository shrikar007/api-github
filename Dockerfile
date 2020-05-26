FROM golang:1.13-alpine AS bob

WORKDIR /go/src/aws-azure-auth-server

RUN apk add --no-cache protobuf git make bash build-base \
	&& rm -rf /var/cache/apk/*

COPY . .

RUN  go build -a -o /main .3

FROM alpine:3.9

WORKDIR /

RUN apk add --no-cache chromium ca-certificates \
	&& update-ca-certificates \
    # cleanup
    && rm -rf /var/cache/apk/*
COPY --from=bob /go/src/github-integration/config.toml .
ENTRYPOINT ["/main"]