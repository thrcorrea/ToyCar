# Base image
FROM golang:alpine

# Set working directory
WORKDIR /app

# Copy the source code
COPY . .

# Build the application
RUN go build -o main -v ./cmd/main.go

RUN chmod +x main

# Expose port
EXPOSE 3000

# Command to run the executable
CMD ["./main"]
