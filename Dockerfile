FROM golang:1.24.1-alpine3.21

WORKDIR /chainUpi

COPY . .

RUN go mod download

RUN go build -o app ./cmd

EXPOSE 3000

CMD ["./app"]