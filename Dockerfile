FROM golang:1.25.6-alpine3.23 AS build

RUN apk add --no-cache git make

WORKDIR /src

RUN git clone https://github.com/weppos/xuuid.git . && \
    CGO_ENABLED=0 GOFLAGS="-ldflags=-s -w" make build

FROM scratch

COPY --from=build /src/xuuid /usr/local/bin/xuuid

ENTRYPOINT ["/usr/local/bin/xuuid"]
