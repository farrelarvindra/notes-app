# Build stage
FROM golang:1.27-alpine AS build
WORKDIR /app

# Salin file dependency (go.sum dibuat opsional dengan wildcard)
COPY go.mod go.sum* ./
RUN go mod download

# Salin source code dan compile binary statis
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -o server main.go

# Run stage
FROM alpine:3.20
WORKDIR /app

# Tambahkan sertifikat SSL & data zona waktu
RUN apk --no-cache add ca-certificates tzdata

COPY --from=build /app/server .
EXPOSE 8080
CMD ["./server"]
