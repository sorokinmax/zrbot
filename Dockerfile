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
RUN apk update && apk upgrade && apk add --no-cache chromium

# Installs latest Chromium package.
RUN echo @edge http://nl.alpinelinux.org/alpine/edge/community >> /etc/apk/repositories \
    && echo @edge http://nl.alpinelinux.org/alpine/edge/main >> /etc/apk/repositories \
    && apk add --no-cache \
    harfbuzz@edge \
    nss@edge \
    freetype@edge \
    ttf-freefont@edge

WORKDIR /zrbot
RUN mkdir -p /zrbot/config
COPY --from=builder /zrbot ./
#RUN chmod +x /zrbot

ENTRYPOINT ["/zrbot/zrbot"]