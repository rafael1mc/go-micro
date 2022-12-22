# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o mailApp ./cmd/api

RUN chmod +x /app/mailApp

# buid a tiny docker image
FROM alpine:latest

RUN mkdir /app

WORKDIR /app

COPY --from=builder /app/mailApp /app
COPY --from=builder /app/templates /app/templates

CMD [ "/app/mailApp" ]