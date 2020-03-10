FROM golang:1.13-alpine as builder

RUN mkdir /app

ADD . /app/

WORKDIR /app

RUN apk add --no-cache git \
        && go get -d -v github.com/gorilla/mux \
	&& go get -d -v gopkg.in/mgo.v2/bson \
	&& go get -d -v gopkg.in/mgo.v2

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .


FROM alpine:latest

EXPOSE 9090

CMD ["/app"]

COPY --from=builder /app/main /app

