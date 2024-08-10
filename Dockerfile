# Start from the official Go image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifest files
COPY go.mod go.sum ./

# Download the Go modules dependencies
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 3000

# Command to run the application
CMD ["./main"]
