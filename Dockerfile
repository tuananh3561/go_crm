ARG GO_VERSION=1.7

#FROM golang:${GO_VERSION}-alpine
FROM golang:alpine

RUN mkdir /app

WORKDIR /app

ADD . .
ADD go.mod .
#ADD go.sum .

RUN go mod tidy -compat=1.17
RUN go mod download
ADD . .

RUN go build -o server cmd/app/main.go

CMD ["./server"]