FROM golang:1.21.1

WORKDIR /searchRecommend/auth 

COPY go.mod /searchRecommend/ 
COPY go.sum /searchRecommend/ 

COPY auth /searchRecommend/auth/ 

RUN go mod tidy