FROM golang:1-alpine

WORKDIR /go/app
COPY . .

RUN go build -o ./bin/main .

CMD [ "bin/main" ]
