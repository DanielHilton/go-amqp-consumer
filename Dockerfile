FROM golang:1.15-alpine

WORKDIR /go/src/github.com/DanielHilton/go-amqp-consumer
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["go-amqp-consumer"]