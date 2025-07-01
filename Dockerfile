# Go API Dockerfile

# Start with a base image containing Go
FROM golang:1.21

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Build the Go app
RUN go mod init stutestalpha && go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
