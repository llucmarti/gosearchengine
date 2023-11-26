
# Get the image
FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy code
COPY . .

# Download and install the dependencies
RUN go get -d -v ./...

RUN go build -o main .

# Expose the port
EXPOSE 8080

# Run the application
CMD ["./main"]
