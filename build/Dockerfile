# Menggunakan image Golang sebagai base image
FROM golang:1.20

# Set working directory
WORKDIR /app

# Menyalin go.mod dan go.sum
COPY go.mod go.sum ./

# Mengunduh dependensi
RUN go mod download

# Menyalin seluruh kode aplikasi
COPY . .

# Mengkompilasi aplikasi
RUN go build -o main .

# Menjalankan aplikasi
CMD ["./main"]
