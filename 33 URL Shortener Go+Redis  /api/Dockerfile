FROM golang:alpine as builder

RUN mkdir /build

ADD . /build

WORKDIR /build

RUN go mod tidy

RUN go build -o shortener-api main.go

FROM alpine

RUN adduser -S -D -H -h/app appuser

USER appuser

COPY . /app

COPY --from=builder /build/shortener-api /app/shortener-api

WORKDIR /app

EXPOSE 3000

CMD ["./shortener-api"]