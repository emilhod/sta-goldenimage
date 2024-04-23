# sta-goldenimage
Image for building and pushing .Net/React to ACR

# Emilhod Golden Image

Emilhod Golden Image is a tool designed to automate the process of building and pushing Docker images for different types of projects to Azure Container Registry. It allows you to easily create a standardized Docker image that can be used across multiple projects.

## Features

- Supports building Docker images for both .NET and React projects.
- Automatically generates dynamic image names based on the repository name, current date, and latest commit hash.
- Provides a simple command-line interface for easy usage.

## Usage

To use Emilhod Golden Image, follow these steps:

1. **Install Docker**: Make sure Docker is installed on your system.
2. **Clone the Repository**: Clone the Emilhod Golden Image repository to your local machine.
3. **Navigate to the Repository Directory**: Open a terminal or command prompt and navigate to the directory where you cloned the repository.
4. **Run the Program**: Execute the program by running the `emilhodgolden` executable file with the appropriate command-line arguments. For example:
   ```bash
   ./emilhodgolden --dotnet --build
   ```
   Replace `--dotnet` with `--react` if you want to build a React project instead.

## Requirements

- Go programming language installed on your system.
- Docker installed on your system.
- Access to an Azure Container Registry (ACR) for pushing Docker images.

## License

This project is licensed under the [MIT License](LICENSE).
