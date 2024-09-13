package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

// UpdateCmd represents the update command
var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update the gg CLI to the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		updateCLI()
	},
}

func updateCLI() {
	repoURL := "https://github.com/Mord1n/gg.git"
	installPath := "/usr/local/bin/gg"

	fmt.Println("Pulling the latest version from GitHub...")

	// Clone the latest version of the repository in a temporary directory
	tmpDir := "/tmp/gg"
	os.RemoveAll(tmpDir) // Clean up any previous tmp directory
	err := exec.Command("git", "clone", repoURL, tmpDir).Run()
	if err != nil {
		fmt.Printf("Error cloning repository: %v\n", err)
		return
	}

	// Change to the directory and build the new version
	fmt.Println("Building the new version...")
	buildCmd := exec.Command("go", "build", "-o", "gg")
	buildCmd.Dir = tmpDir
	err = buildCmd.Run()
	if err != nil {
		fmt.Printf("Error building the new version: %v\n", err)
		return
	}

	// Replace the old binary with the new one
	fmt.Println("Replacing the old version with the new one...")
	err = exec.Command("sudo", "mv", tmpDir+"/gg", installPath).Run()
	if err != nil {
		fmt.Printf("Error updating binary: %v\n", err)
		return
	}

	fmt.Println("gg CLI updated successfully!")
}
