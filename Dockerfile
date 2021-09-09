FROM golang:1.16-alpine3.14

WORKDIR /src/

RUN mkdir api.lyabook

WORKDIR /src/api.lyabook

ENV BASE_URI = '/api/v1'
ENV URL_MONGO='mongodb+srv://lyabook:morsac@lyabook.d1ttq.mongodb.net/lyabook?retryWrites=true&w=majority'
ENV DB_NAME='lyabook'
ENV SECRET_KEY='eyJ0eXAicmcx100preOiJKV1QiLCJhbGciOiJIUzI1NiJ9'
ENV PORT=8000
ENV EMAIL='leo2001.nl08@gmail.com'
ENV PASSWORD_EMAIL='gxyt rxyx vimx yshl'
ENV API_RENIEC_DNI='https://api.apis.net.pe/v1/dni?numero='
ENV TOKEN_API_RENIEC_DNI='apis-token-1.aTSI1U7KEuT-6bbbCguH-4Y8TI6KS73N'

COPY . .

RUN go build -o app src/main.go

EXPOSE 8000

CMD ["./app"]

