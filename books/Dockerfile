FROM golang:1.21.1

WORKDIR /searchRecommend/books 

COPY go.mod /searchRecommend/ 
COPY go.sum /searchRecommend/ 
COPY credentials.json /searchRecommend/ 

COPY books /searchRecommend/books/ 

RUN go mod tidy