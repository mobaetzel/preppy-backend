FROM golang:1.19-alpine AS build

WORKDIR /code

COPY ./go.mod ./
COPY ./go.sum ./
COPY ./src ./src

RUN go mod download

RUN go build -o /app ./src/

EXPOSE 8000

ENTRYPOINT [ "/app" ]