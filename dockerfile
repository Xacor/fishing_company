FROM golang:1.19 AS builder
WORKDIR /usr/src/
RUN git clone https://github.com/Xacor/fishing_company.git
WORKDIR /usr/src/fishing_company
RUN RUN go mod download && go mod verify
RUN go build -v -o app cmd/main.go



FROM alpine:3.16
WORKDIR /root/
COPY --from=builder /usr/src/fishing_company/app ./
COPY --from=builder /usr/src/fishing_company/ui ./
CMD ["./app"]