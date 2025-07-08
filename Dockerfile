# Dockerfile (exemplo básico para app Go)
FROM golang:1.21 AS builder
WORKDIR /app
COPY backend .
RUN go build -o app ./cmd

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]
