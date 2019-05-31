package utils

import (
	"os"
	"os/exec"
	"log"
)

// RunScript takes a script path as a string and executes it while
// sending stdout/stderr to os default
func RunScript(script string) {
	cmd := exec.Command("sh", script)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Test harness failed to run with error: %s\n", err)
	}
}