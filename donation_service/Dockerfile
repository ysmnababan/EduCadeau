# Menggunakan image golang sebagai base image
FROM golang:1.22-alpine

# Mengatur working directory
WORKDIR /app

# Mengcopy go.mod dan go.sum lalu install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Mengcopy kode sumber ke dalam container
COPY . .

# Build aplikasi
RUN go build -o main .

# Expose port 50001 to the outside world
EXPOSE 50001

# Menjalankan aplikasi
CMD ["./main"]