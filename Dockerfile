FROM golang:1.10

RUN mkdir -p /go/src/app
RUN go get -u github.com/gorilla/mux
WORKDIR /go/src/app

ADD ./server /go/src/app

RUN go get -v