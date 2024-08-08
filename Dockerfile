FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY . .
RUN go mod download

RUN go build -o ./url-shortener

FROM scratch

WORKDIR /app
COPY --from=builder /build/url-shortener ./url-shortener

CMD ["/app/url-shortener"]
