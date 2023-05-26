# build stage
FROM golang:1.20.4-alpine3.17 AS builder
ADD . /src
RUN apk add git
RUN cd /src && go build -o /src/bin/api -v /src/cmd/

# final stage
FROM alpine:3.17

WORKDIR /app
COPY --from=builder /src/bin/api /tmp/api
ENTRYPOINT /tmp/api