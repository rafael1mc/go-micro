# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o frontendApp ./cmd/web

RUN chmod +x frontendApp

# buid a tiny docker image
FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app /app

CMD [ "./frontendApp" ]
