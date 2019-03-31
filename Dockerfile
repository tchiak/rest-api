FROM golang:1.12

WORKDIR /go/src/rest-api
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["rest-api"]