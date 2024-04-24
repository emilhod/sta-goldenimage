# Use an official base image with Go installed
FROM golang:1.22

# Install jq and yq
RUN apt-get update && \
    apt-get install -y jq && \
    apt-get install -y yq && \
    rm -rf /var/lib/apt/lists/*

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o emilhodgolden .

# Command to run the Go application
CMD ["./emilhodgolden"]
