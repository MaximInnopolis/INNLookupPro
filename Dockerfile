# Use the official Go image as a parent image
FROM golang:1.17 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules manifests
COPY go.mod go.sum ./

# Download the Go modules
RUN go mod download

# Copy the source code from your local machine to the container
COPY . .

# Build the application
RUN go build -o app

# Start a new stage from a minimal image
FROM gcr.io/distroless/base

# Set the working directory inside the container
WORKDIR /app

# Copy the built application from the previous stage
COPY --from=builder /app/app .

# Expose the port the service will run on
EXPOSE 8080

# Run the application
CMD ["./app"]
