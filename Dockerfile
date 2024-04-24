# Use an official base image with Go installed
FROM golang:1.22

# Install jq and yq
RUN apt-get update && \
    apt-get install -y jq && \
    apt-get install -y yq && \
    rm -rf /var/lib/apt/lists/*

# Install kubeconform
RUN go install github.com/yannh/kubeconform/cmd/kubeconform@latest
RUN go install golang.stackrox.io/kube-linter/cmd/kube-linter@latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

# Build the Go application
RUN go build -o emilhodgolden .

# Command to run the Go application
CMD ["./emilhodgolden"]
