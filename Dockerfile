##
## Build stage
##
FROM golang:1.19.2-alpine3.16 AS builder
RUN apk update && apk upgrade
WORKDIR /src
COPY . .

RUN go mod download
RUN GOOS=linux go build -o /zrbot .

##
## Final image stage
##
FROM alpine:3.16
WORKDIR /zrbot
RUN mkdir config
COPY --from=builder /zrbot /zrbot

ENTRYPOINT ["/zrbot/zrbot"]