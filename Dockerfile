# Use an official base image with Go installed
FROM golang:1.22

# Install jq and yq
RUN apt-get update && \
    apt-get install -y jq && \
    apt-get install -y yq && \
    rm -rf /var/lib/apt/lists/*

# Install Trivy
RUN apt-get install wget apt-transport-https gnupg lsb-release \
    wget -qO - https://aquasecurity.github.io/trivy-repo/deb/public.key | gpg --dearmor | tee /usr/share/keyrings/trivy.gpg > /dev/null \
    echo "deb [signed-by=/usr/share/keyrings/trivy.gpg] https://aquasecurity.github.io/trivy-repo/deb $(lsb_release -sc) main" | tee -a /etc/apt/sources.list.d/trivy.list \
    apt-get update \
    apt-get install trivy \
    rm -rf /var/lib/apt/lists/*


# Install kubeconform & kubelinter
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
