FROM golang:1.19

WORKDIR $GOPATH/src/github.com/arulselvan/wizchat/router-service

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["router-service"]