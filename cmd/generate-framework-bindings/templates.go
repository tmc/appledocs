package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"

	"golang.org/x/tools/txtar"
)

var templateArchive *txtar.Archive
var variantArchives map[string]*txtar.Archive

func init() {
	// Load configuration
	if err := loadConfig(); err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	// Load base templates
	templatesData, err := embeddedFS.ReadFile("templates.txtar")
	if err != nil {
		panic(fmt.Errorf("failed to read templates.txtar: %w", err))
	}
	templateArchive = txtar.Parse(templatesData)

	// Load variant templates
	variantArchives = make(map[string]*txtar.Archive)
	entries, err := fs.ReadDir(embeddedFS, ".")
	if err == nil {
		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}
			name := entry.Name()
			// Match templates_<variant>.txtar pattern
			if strings.HasPrefix(name, "templates_") && strings.HasSuffix(name, ".txtar") {
				variantName := strings.TrimSuffix(strings.TrimPrefix(name, "templates_"), ".txtar")
				data, err := embeddedFS.ReadFile(name)
				if err == nil {
					variantArchives[variantName] = txtar.Parse(data)
					if verbose {
						fmt.Fprintf(os.Stderr, "Loaded template variant: %s\n", variantName)
					}
				}
			}
		}
	}

}

// getTemplateVariant returns the template content for a given file, checking variant archives.
// Supports comma-separated variant layering from separate files. For example:
//
//	variant="darwinkit,ref-methods" and filename="types.gen.go" will look for (in order):
//	1. types.gen.go in templates_ref-methods.txtar
//	2. types.gen.go in templates_darwinkit.txtar
//	3. types.gen.go in templates.txtar (fallback)
//
// The first match wins, allowing later variants to override earlier ones.
func getTemplateVariant(filename, variant string) (string, error) {
	// Try variants in reverse order (rightmost first) so later variants override earlier ones
	if variant != "" {
		variants := strings.Split(variant, ",")
		for i := len(variants) - 1; i >= 0; i-- {
			v := strings.TrimSpace(variants[i])
			if v == "" {
				continue
			}
			// Look in variant archive
			if archive, ok := variantArchives[v]; ok {
				for _, file := range archive.Files {
					if file.Name == filename {
						return string(file.Data), nil
					}
				}
			}
		}
	}

	// Fallback to base template
	for _, file := range templateArchive.Files {
		if file.Name == filename {
			return string(file.Data), nil
		}
	}

	return "", fmt.Errorf("template not found: %s", filename)
}
