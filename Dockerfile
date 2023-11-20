
# Step 1: Choose a base image
FROM golang:latest

# Step 2: Set the working directory
WORKDIR /app

# Step 3: Copy your code
COPY . .

# Download and install the dependencies
RUN go get -d -v ./...

RUN go build -o main .

#EXPOSE the port
EXPOSE 8000

# Step 6: Set the command to run your code
CMD ["./main"]
