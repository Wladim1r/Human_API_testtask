# Build
FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/bin/testtask ./cmd/main.go

# Final
FROM alpine:3.21

WORKDIR /api

COPY --from=builder /app/bin/testtask .
COPY --from=builder /app/internal/db /api/internal/db

EXPOSE 8080

CMD [ "./testtask" ]
