FROM golang:1.23-alpine as builder

# Set directory
WORKDIR /app

# Download dependencies and libs
COPY go.mod go.sum ./
RUN go mod download

# Copy service to current directory
COPY . .

# Set directory to role service
WORKDIR /app/services/role

# Build service
RUN go build -o role /cmd/root.go

# Command
CMD ["./role"]

