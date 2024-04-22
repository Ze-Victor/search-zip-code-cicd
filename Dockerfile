FROM golang:1.21.1

ADD . /app

WORKDIR /app/cmd/api

RUN go build -o main .

EXPOSE 8001

CMD ["./main"]
