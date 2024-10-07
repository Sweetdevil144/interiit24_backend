# Build stage: Use Go 1.22 to compile the application
FROM golang:1.22-alpine AS build

# Set the working directory inside the container
WORKDIR /app

# Copy Go modules files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the application binary
RUN go build -o compute_server .

# Final stage: Use a minimal Alpine image to run the binary
FROM alpine:3.17

# Set the working directory
WORKDIR /app

# Copy the binary from the build stage
COPY --from=build /app/compute_server /app/compute_server

# Expose the necessary ports (6942, 6943, 6944)
EXPOSE 6942 6943 6944

# Command to run the compute server on different ports
CMD ["sh", "-c", "PORT=$PORT /app/compute_server"]
