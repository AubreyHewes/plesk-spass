FROM golang:1.13-alpine3.10 as builder

RUN apk --no-cache --no-progress add make git upx

WORKDIR /go/src/github.com/aubreyhewes/plesk-spass
COPY . .
RUN make build

FROM alpine:3.10
RUN apk update \
    && apk add --no-cache ca-certificates tzdata \
    && update-ca-certificates

COPY --from=builder /go/src/github.com/aubreyhewes/plesk-spass/dist/spass /usr/bin/spass
ENTRYPOINT [ "/usr/bin/spass" ]