FROM golang:1.23-alpine as builder

# Set directory
WORKDIR /app

# Download dependencies and libs
COPY go.mod go.sum ./
RUN go mod download

# Copy service to current directory
COPY . .

# Set directory to user service
WORKDIR /app/services/user

# Build service
RUN go build -o user main.go

# Command
CMD ["./user"]