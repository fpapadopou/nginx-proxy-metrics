# Start from the latest Go image
FROM golang:1.13

# Set the working directory
WORKDIR /app

# Copy source code to container
COPY . .

# Build the Go app
RUN go build -o ./cmd/greet/main ./cmd/greet

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
ENTRYPOINT ["./cmd/greet/main"]