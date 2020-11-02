FROM golang:1.14.0

WORKDIR /app

RUN go build -o main .

COPY main /app
ENTRYPOINT [./app/main]