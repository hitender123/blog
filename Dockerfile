# Start from the official Golang image
FROM golang:1.19

# Set the working directory inside the container
WORKDIR /go/src/blog

# Copy the application files to the container
COPY . .

# Install the application dependencies
RUN go mod download

# Build the application
RUN go build -o blog .

# Expose the application port
EXPOSE 8080

# Set the command to run when the container starts
CMD ["./blog"]
