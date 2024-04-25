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
	build := buildDockerImage(projectType, imageName)
	if build != nil {
		fmt.Println(build) //print error
		return
	}

	//Docker tag
	tag := tagDockerImage(projectType, imageName, fullImageName)
	if tag != nil {
		fmt.Println(tag) //print error
		return
	}

	//Docker push
	push := pushDockerImage(fullImageName)
	if push != nil {
		fmt.Println(push) //print error
		return
	}

	//If project is go, also tag and push latest tag
	if projectType == "Golang"{
		fmt.Println("Also adding latest tag...")
		repoName := getRepositoryName()
		fullLatestImageName := acrName + "/" + repoName + ":latest"

		tag := tagDockerImage(projectType, imageName, fullLatestImageName)
		if tag != nil {
			fmt.Println(tag) //print error
			return
		}

		push := pushDockerImage(fullLatestImageName)
		if push != nil {
			fmt.Println(push) //print error
			return
		}
	}
}


// generateImageName generates a dynamic image name based on the repository name, current date, and latest commit hash
func generateImageName() string {
	// Get repository name
	repoName := getRepositoryName()

	// Get current date
	currentDate := time.Now().Format("02-01-2006")

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

func getRepositoryName() string {
	// Get repository name
	repoNameCmd := exec.Command("git", "rev-parse", "--show-toplevel")
	repoNameOutput, err := repoNameCmd.Output()
	if err != nil {
		fmt.Println("Error getting repository name:", err)
		return ""
	}
	repoName := filepath.Base(strings.TrimSpace(string(repoNameOutput)))
	return repoName
}

func buildDockerImage(projectType, imageName string) error {
	fmt.Printf("Building Docker image for %s project with name: %s...\n", projectType, imageName)

	cmd := exec.Command("docker", "build", "-t", imageName, ".")
	cmd.Dir = "./" // Set the working directory to the root of the repository
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error building Docker image for %s project: %s\n%s", projectType, err, string(output))
	}

	fmt.Println("Docker image built successfully!")
	return nil
}

func tagDockerImage(projectType, imageName, fullImageName string) error {
	fmt.Printf("Tagging the image...\n")
	tagCmd := exec.Command("docker", "tag", imageName, fullImageName)
	tagOutput, err := tagCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error tagging image for %s project: %s\n%s", projectType, err, string(tagOutput))
	}
	fmt.Println("Docker image tagged successfully!")
	return nil
}

func pushDockerImage(fullImageName string) error {
		// Push the Docker image to Azure Container Registry
	fmt.Printf("Pushing Docker image to Azure Container Registry...\n")
	pushCmd := exec.Command("docker", "push", fullImageName)
	pushOutput, err := pushCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error pushing Docker image to Azure Container Registry: %s\n%s", err, string(pushOutput))
		
	}
	fmt.Println("Docker image pushed to Azure Container Registry successfully!")
	return nil
}