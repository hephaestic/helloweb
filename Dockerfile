FROM golang:latest

WORKDIR /go/src

ADD . /go/src/helloweb

EXPOSE 8080

RUN go install helloweb

CMD ["helloweb"]

