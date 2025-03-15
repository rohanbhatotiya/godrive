package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
)

const version = "1.0.0"

const toolInfo = `
Godrive CLI Tool
----------------
A command-line tool to easily upload files and directories to Google Drive.

Features:
- First-time user setup for Google authentication
- Supports file and directory uploads
- Resumes authentication if previously authenticated

Developed by Rohan Yadav
LinkedIn -> https://www.linkedin.com/in/rohanbhatotiya
GitHub -> https://github.com/rohanbhatotiya
Email -> rohanbhatotiya@gmail.com
`

func promptUser(prompt string) string {
	fmt.Print(prompt)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func verifyClientSecrets(content string) bool {
	return len(content) > 0 && len(content) < 5000 && content[0] == '{' && content[len(content)-1] == '}'
}

func getClientSecrets(clientSecretsPath string) {
	attempts := 0
	for attempts < 3 {
		clientSecrets := promptUser("\nPaste your client_secrets.json content:\n")
		fmt.Println("Verifying...")
		if verifyClientSecrets(clientSecrets) {
			err := os.WriteFile(clientSecretsPath, []byte(clientSecrets), 0644)
			if err != nil {
				fmt.Println(" Error saving client_secrets.json:", err)
				return
			}
			fmt.Println("✅ Successfully verified and saved.")
			return
		} else {
			fmt.Println(" Invalid client_secrets.json format. Please try again.")
			attempts++
		}
	}
	fmt.Println("Too many failed attempts. Exiting...")
	os.Exit(1)
}

func verifyFilePath(filePath string) bool {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		fmt.Println(" Error resolving absolute path:", err)
		return false
	}

	if _, err := os.Stat(absPath); err != nil {
		fmt.Println(" File/Directory does not exist:", err)
		return false
	}

	return true
}

func printVersion() {
	fmt.Println("Godrive CLI Version:", version)
}

func printInfo() {
	fmt.Println(toolInfo)
}

func main() {
	// Handle interrupt signals (Ctrl+C)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		fmt.Println("\n Exiting gracefully... Goodbye!")
		os.Exit(0)
	}()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(" Error: Unable to determine home directory.")
		os.Exit(1)
	}

	godriveDir := filepath.Join(homeDir, ".godrive")
	clientSecretsPath := filepath.Join(godriveDir, "client_secrets.json")
	authFilePath := filepath.Join(godriveDir, "auth.txt")

	// Ensure .godrive directory exists
	if err := os.MkdirAll(godriveDir, 0755); err != nil {
		fmt.Println(" Error creating .godrive directory:", err)
		os.Exit(1)
	}

	// Check for CLI commands
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "version":
			printVersion()
			return
		case "info":
			printInfo()
			return
		}
	}

	fmt.Println("🚀 Welcome to Godrive CLI!")

	// Check client_secrets.json existence and validity
	fileContent, err := os.ReadFile(clientSecretsPath)
	if err != nil || !verifyClientSecrets(string(fileContent)) {
		fmt.Println("\n⚠️  Missing or invalid client_secrets.json. You need to provide a new one.")
		getClientSecrets(clientSecretsPath)
	} else {
		choice := promptUser("\n✅ Valid client_secrets.json found. Do you want to change it? (y/n): ")
		if choice == "y" {
			getClientSecrets(clientSecretsPath)
		} else {
			fmt.Println("✅ Using existing client_secrets.json.")
		}
	}

	// Get the file/directory to upload
	var filePath string
	attempts := 0

	for {
		fmt.Print("\n Enter the full path of the file/directory to upload: ")
		fmt.Scanln(&filePath)

		if filePath == "" {
			fmt.Println(" Error: Path cannot be empty. Please enter a valid path.")
		} else if _, err := os.Stat(filePath); os.IsNotExist(err) {
			fmt.Println(" Error: The specified file/directory does not exist. Try again.")
			attempts++
		} else {
			break
		}

		if attempts >= 3 {
			fmt.Println(" Too many invalid attempts. Exiting...")
			os.Exit(1)
		}
	}

	// Get an optional name for the uploaded file
	fileName := promptUser("\n Enter an optional name for the uploaded file (Press enter to keep original): ")

	// Check authentication status
	if _, err := os.Stat(authFilePath); err == nil {
		prevUser := promptUser("\n Previous upload was done on a saved account. Do you want to use the same account? (y/n): ")
		if prevUser == "n" {
			os.Remove(authFilePath)
			fmt.Println(" Old authentication removed. You will need to authenticate again.")
		} else {
			fmt.Println("✅ Using existing authentication.")
		}
	}

	// Find upload executable
	exePath, err := os.Executable()
	if err != nil {
    	fmt.Println("❌ Error: Unable to determine executable path.")
    	os.Exit(1)
	}

	// Get the directory of the executable
	exeDir := filepath.Dir(exePath)

	// Construct the path to the upload executable
	uploadExec := filepath.Join(exeDir, "godrive_upload")
	fmt.Println("🔍 Looking for godrive_upload at:", uploadExec)

	// Check if godrive_upload exists
	if _, err := os.Stat(uploadExec); os.IsNotExist(err) {
    	fmt.Println("❌ Error: The upload script is missing. Please reinstall Godrive.")
    	os.Exit(1)
	}

	// Run godrive_upload with arguments
	cmd := exec.Command(uploadExec, filePath, fileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		fmt.Println(" Upload failed! Ensure Python is installed and retry.")
		os.Exit(1)
	}
}
