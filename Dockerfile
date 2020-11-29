# build
FROM golang:1.15.5-alpine3.12 as build

ENV PORT 8080
EXPOSE 8080

RUN mkdir /app
ADD . /app

ENV GOPROXY https://goproxy.io
ENV GIN_MODE release

WORKDIR  /app
RUN go mod vendor
RUN go build -mod=vendor -o golang-logger .


# release
FROM alpine:3.12
RUN mkdir /app
COPY --from=build /app/golang-logger /app/golang-logger

WORKDIR  /app
CMD ["/app/golang-logger"]
