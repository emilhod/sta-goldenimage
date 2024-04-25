# sta-goldenimage
Image for building and pushing .Net/React to ACR

# Golden/Base Image

This Golden/Base Image is a tool designed to automate the process of building and pushing Docker images for different types of projects to Azure Container Registry. It allows you to easily create a standardized Docker image that can be used across multiple projects.

## Features

- Supports building Docker images for .NET, React and Go projects.
- Automatically generates dynamic image names based on the repository name, current date, and latest commit hash.
- Provides a simple command-line interface for easy usage.
- Includes pre-installed tools for enhanced functionality:
    - **jq and yq**: Command-line tools for parsing and manipulating JSON and YAML data.
    - **Trivy**: A comprehensive vulnerability scanner for containers, providing insights into security issues within container images.
    - **kubeconform & kubelinter**: Tools for Kubernetes configuration validation and linting, ensuring compliance with best practices and standards.

## Usage

To use image, follow these steps:

1. **Install Docker**: Make sure Docker is installed on your system.
2. **Clone the Repository**: Clone the Emilhod Golden Image repository to your local machine.
3. **Navigate to the Repository Directory**: Open a terminal or command prompt and navigate to the directory where you cloned the repository.
4. **Run the Program**: Execute the program by running the `emilhodgolden` executable file with the appropriate command-line arguments. For example:
   ```bash
   ./emilhodgolden --dotnet --build
   ```
   Replace `--dotnet` with `--react`  or `--go` if you want to build another project instead.

## Requirements

- Go programming language installed on your system.
- Docker installed on your system.
- Access to an Azure Container Registry (ACR) for pushing Docker images.

## License

This project is licensed under the [MIT License](LICENSE).
