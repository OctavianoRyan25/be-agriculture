#Build stage
FROM golang:1.22-alpine3.19 AS build-stage

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o /goapp

#Release stage
FROM alpine:3.19 AS build-release-stage

WORKDIR /

COPY --from=build-stage /goapp /goapp
COPY --from=build-stage /app/modules /app/modules

EXPOSE 8080

ENTRYPOINT ["./goapp"]