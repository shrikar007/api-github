FROM golang:1.13-alpine AS bob

WORKDIR /go/src/github-integration

RUN apk add --no-cache git bash build-base \
	&& rm -rf /var/cache/apk/*

COPY . .

RUN  go build -a -o /main .

FROM alpine:3.9

WORKDIR /

RUN apk add --no-cache chromium ca-certificates \
	&& update-ca-certificates \
    # cleanup
    && rm -rf /var/cache/apk/*
COPY --from=bob /main .
COPY --from=bob /go/src/github-integration/config.toml .
ENTRYPOINT ["/main"]
