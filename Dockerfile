FROM golang:1.17.1-alpine3.14 as builder

RUN mkdir /build

ADD . /build
WORKDIR /build
RUN go build -race -o runner .

FROM scratch
COPY --from=builder /build/runner /app/
WORKDIR /app

CMD [ "./runner" ]
