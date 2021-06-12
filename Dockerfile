# builder
FROM golang:1.16.5-alpine3.13

RUN apk update && \
    apk upgrade && \
    apk add make

WORKDIR /app

COPY . .

CMD ["make", "run"]