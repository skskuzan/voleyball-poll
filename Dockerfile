FROM golang:1.19-alpine

WORKDIR /app

COPY . .

RUN go build -o volleyball_bot

CMD ["./volleyball_bot"]
