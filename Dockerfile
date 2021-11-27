FROM golang:1.17.2

RUN apt update && apt install -y wkhtmltopdf

WORKDIR /src/

RUN mkdir api.lyabook

WORKDIR /src/api.lyabook

COPY . .

ENV BASE_URI=${BASE_URI}
ENV APP_CLIENT=${APP_CLIENT}
ENV URL_MONGO=${URL_MONGO}
ENV DB_NAME=${DB_NAME}
ENV SECRET_KEY=${SECRET_KEY}
ENV PORT=${PORT}
ENV EMAIL=${EMAIL}
ENV PASSWORD_EMAIL=${PASSWORD_EMAIL}
ENV API_RENIEC_DNI=${API_RENIEC_DNI}
ENV TOKEN_API_RENIEC_DNI=${TOKEN_API_RENIEC_DNI}

RUN go build -o app src/main.go

EXPOSE 8000

CMD ["./app"]
