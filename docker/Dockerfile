FROM golang:1.22.4-alpine as builder
WORKDIR /bank

COPY . .

RUN go mod download

RUN GOARCH=arm64 CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/*
RUN chmod +x app

FROM scratch as runner

COPY --from=builder /bank/app .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

EXPOSE 3000

USER nobody:nogroup

CMD ["/app"]