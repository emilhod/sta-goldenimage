package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	args := os.Args[1:]


	// Check if the provided flag is --dotnet, --go or --react
	if len(args) < 1 || (args[0] != "--dotnet" && args[0] != "--react" && args[0] != "--go") {
    	fmt.Println("Invalid command. Usage: emilhodgolden (--dotnet | --react | --go) --build")
    	return
	}

	// Determine the project type based on the provided flag
	var projectType string
	switch args[0] {
	case "--dotnet":
		projectType = ".NET"
	case "--react":
		projectType = "React"
	case "--go":
		projectType = "Golang"
	default:
		fmt.Println("Invalid command. Usage: emilhodgolden (--dotnet | --react | --go) --build")
		return
	}

	// Generate the image name
	imageName := generateImageName()
	acrName := os.Getenv("AZURE_ACR_NAME_DEV")
	fullImageName := acrName + "/" + imageName

	// Execute the Docker build command for the project
	fmt.Printf("Building Docker image for %s project with name: %s...\n", projectType, imageName)

	cmd := exec.Command("docker", "build", "-t", imageName, ".")
	cmd.Dir = "./" // Set the working directory to the root of the repository
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error building Docker image for %s project: %s\n", projectType, err)
		fmt.Println(string(output))
		return
	}

	fmt.Printf("Docker image for %s project built successfully! The image name is %s\n", projectType, imageName)


	// Push the Docker image to Azure Container Registry
	fmt.Printf("Pushing Docker image to Azure Container Registry...\n")

	
	fmt.Println("Acr name is :", acrName)
	pushCmd := exec.Command("docker", "push", fullImageName)
	pushOutput, err := pushCmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error pushing Docker image to Azure Container Registry: %s\n", err)
		fmt.Println(string(pushOutput))
		return
	}

	fmt.Println("Docker image pushed to Azure Container Registry successfully!")
	
}



// generateImageName generates a dynamic image name based on the repository name, current date, and latest commit hash
func generateImageName() string {
	// Get repository name
	repoNameCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	repoNameOutput, err := repoNameCmd.Output()
	if err != nil {
		fmt.Println("Error getting repository name:", err)
		return ""
	}
	repoName := filepath.Base(strings.TrimSpace(string(repoNameOutput)))

	// Get current date
	currentDate := time.Now().Format("2006-01-02")

	// Get latest commit hash
	commitHashCmd := exec.Command("git", "rev-parse", "--short=7", "HEAD")
	commitHashOutput, err := commitHashCmd.Output()
	if err != nil {
		fmt.Println("Error getting latest commit hash:", err)
		return ""
	}
	commitHash := strings.TrimSpace(string(commitHashOutput))

	// Construct the image name
	imageName := fmt.Sprintf("%s:%s-%s", repoName, currentDate, commitHash)
	return imageName
}
