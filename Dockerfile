FROM golang:1.16-alpine3.14

WORKDIR /src/

RUN mkdir api.lyabook

WORKDIR /src/api.lyabook

COPY . .

RUN go build -o app src/main.go

CMD ["./app"]
