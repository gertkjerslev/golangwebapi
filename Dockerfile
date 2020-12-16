FROM golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/gertkjerslev/golangwebapi/webapi/src
RUN cd /build && git clone https://github.com/gertkjerslev/golangwebapi.git

RUN cd /build/golangwebapi/webapi/src && go build

EXPOSE 8080

ENTRYPOINT ["/build/golangwebapi/webapi/src/main"]