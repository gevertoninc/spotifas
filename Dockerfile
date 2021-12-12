# syntax=docker/dockerfile:1
FROM golang:1.17.3-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o /spotifas

FROM golang:1.17.3-alpine

WORKDIR /app

COPY --from=build /spotifas /app/spotifas

EXPOSE 8080

CMD ["/spotifas"]
