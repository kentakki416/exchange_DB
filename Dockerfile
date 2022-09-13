FROM golang:1.17.7-alpine
 
RUN apk update && \
    apk upgrade && \
    apk add git
 
RUN go install github.com/cespare/reflex@latest
RUN go get github.com/guregu/dynamo
ENV CGO_ENABLED=0
 
WORKDIR /go/src/app
COPY go.* main.go ./
 
RUN go mod download