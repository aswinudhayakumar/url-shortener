FROM golang:1.21-alpine

WORKDIR /app
COPY go.mod go.sum ./

RUN go mod download

COPY main.go ./
COPY helper/ ./helper
COPY model/ ./model
COPY router/ ./router

RUN CGO_ENABLED=0 GOOS=linux go build -o /url-shortener

EXPOSE 3000

CMD ["/url-shortener"]
