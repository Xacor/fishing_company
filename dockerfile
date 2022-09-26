FROM golang:1.19 AS builder
WORKDIR /usr/src/fishing_company
COPY . .
RUN go build -v -o app cmd/main.go
EXPOSE 5000
CMD ["./app"]