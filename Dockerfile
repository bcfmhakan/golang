FROM golang:1.12.0-alpine3.9
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN apk add git
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/etherlabsio/healthcheck/v2
RUN go get github.com/etherlabsio/healthcheck/v2/checkers
RUN go build -o main .
CMD ["/app/main"]
