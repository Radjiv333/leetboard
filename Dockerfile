# Use official Go image to build the app
FROM golang:1.25.6

# Create working directory
WORKDIR /app

# Copy go.mod and go.sum first (better caching)
COPY go.mod ./

# Download dependencies
RUN go mod download

# Copy the rest of the project
COPY . .

# Build the application
RUN go build -o leetboard cmd/main.go

# Expose port (optional)
EXPOSE 8080

# Run the program
CMD ["./leetboard"]