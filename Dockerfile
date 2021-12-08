FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk install git
RUN go get github.com/prometheus/client_golang/prometheus
RUN go build -o main .
CMD ["/app/main"]
