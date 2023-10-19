FROM golang:1.21.1

WORKDIR /go/src/

COPY go.mod /go/src/ 
COPY go.sum /go/src/ 
COPY credentials.json /go/src/

RUN go mod tidy

COPY . /go/src/

EXPOSE 9090