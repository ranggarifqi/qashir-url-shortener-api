# builder
FROM golang:1.16.5-alpine3.13

WORKDIR /app

COPY . .

RUN go build -o api cmd/api/main.go

CMD /app/api