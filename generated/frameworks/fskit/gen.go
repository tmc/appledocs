//go:build ignore

package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	// Remove generated files
	generatedFiles := []string{
		"doc.go",
		"types.gen.go",
		"functions.gen.go",
		"loader.gen.go",
	}

	for _, file := range generatedFiles {
		if err := os.Remove(file); err != nil && !os.IsNotExist(err) {
			log.Printf("Warning: failed to remove %s: %v", file, err)
		}
	}

	// Re-generate bindings using the appledocs command from PATH
	cmd := exec.Command("appledocs", "generate-framework", "FSKit")

	// Set working directory to project root
	// We're in generated/frameworks/fskit, so go up 3 levels
	if wd, err := os.Getwd(); err == nil {
		projectRoot := filepath.Join(wd, "..", "..", "..")
		cmd.Dir = projectRoot
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatalf("Failed to regenerate FSKit bindings: %v", err)
	}

	log.Println("Successfully regenerated FSKit bindings")
}
