FROM golang:1.21.5-alpine3.18

WORKDIR /app

COPY . .

RUN go build -o api

EXPOSE 8080

CMD ["./api"]