FROM golang:1.17.5-alpine3.15 as build

WORKDIR /opt/app
COPY . .

RUN go mod tidy && go build main.go

FROM alpine

WORKDIR /app
COPY --from=build /opt/app/main /app/main
COPY --from=build /opt/app/static /app/static

CMD [ "./main" ]
