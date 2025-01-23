FROM golang:1.17.1-alpine3.14 as builder

RUN mkdir /build

ADD . /build
WORKDIR /build
RUN apk update  \
    && apk add --virtual  \
    build-dependencies  \
    build-base  \
    gcc  \
    && go mod vendor  \
    && go build -race -o runner .

FROM alpine:3.14.0
COPY --from=builder /bin/runner /app/
WORKDIR /app

CMD [ "./runner" ]
