# Menggunakan image Alpine Linux dengan Golang
FROM golang:1.23.1-alpine

# Set working directory
WORKDIR /app

# Menyalin go.mod dan go.sum
COPY go.mod go.sum ./

# Mengunduh dependensi
RUN go mod download

# Menyalin seluruh kode aplikasi
COPY . .

# Mengkompilasi aplikasi dengan CGO dinonaktifkan (untuk Alpine)
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/gopatungan

# Menjalankan aplikasi
CMD ["./app"]