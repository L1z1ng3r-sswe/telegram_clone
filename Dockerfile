# Use a Golang base image
FROM golang:latest AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules and build files to the working directory
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code from the current directory to the working directory inside the container
COPY . .

# Copy the config directory
COPY config/ ./config

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o ./app/cmd/telegram_clone/build/binary_app ./app/cmd/telegram_clone/main.go

# Start a new stage from scratch
FROM alpine:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the binary built in the previous stage
COPY --from=builder /app/app/cmd/telegram_clone/build/binary_app .

# Copy the config directory
COPY --from=builder /app/config ./config

# Expose any ports the app needs
EXPOSE 8080

# Command to run the executable
CMD ["./binary_app"]
