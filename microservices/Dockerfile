FROM golang:alpine

ADD . /go/src/gin_rummy
# RUN go install gin_rummy
WORKDIR /go/src/gin_rummy
RUN go get .

ENTRYPOINT /go/bin/gin_rummy 

EXPOSE 9090