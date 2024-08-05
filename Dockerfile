FROM golang:1.21
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o cmd/server/main ./cmd/server/main.go
COPY .env .env
EXPOSE 8080
CMD ["./cmd/server/main"]