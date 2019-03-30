FROM golang:1.8

WORKDIR /go/src/golang-webapi
COPY go/src/golang-webapi /go/src/golang-webapi

RUN go get -v github.com/beego/bee
RUN go get -v .
RUN go get -d -v ./...
# RUN go install -v golang-webapi
expose 8888

CMD ["bee", "run", "-downdoc=true", "-gendoc=true"]

