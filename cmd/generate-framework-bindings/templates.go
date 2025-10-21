package main

import (
	"fmt"
	"io/fs"
	"os"
	"strings"
	"text/template"

	"golang.org/x/tools/txtar"
)

var (
	docTemplate               *template.Template
	coreGraphicsTypesTemplate *template.Template
	classesTemplate           *template.Template
	protocolsTemplate         *template.Template
	functionsGenTemplate      *template.Template
	methodsTemplate           *template.Template
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

	templates := make(map[string]string)
	for _, file := range templateArchive.Files {
		templates[file.Name] = string(file.Data)
	}

	var parseErr error
	docTemplate, parseErr = template.New("doc.gen.go").Funcs(templateFuncs).Parse(templates["doc.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse doc.gen.go: %w", parseErr))
	}

	coreGraphicsTypesTemplate, parseErr = template.New("types.gen.go").Funcs(templateFuncs).Parse(templates["types.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse types.gen.go: %w", parseErr))
	}

	classesTemplate, parseErr = template.New("classes.gen.go").Funcs(templateFuncs).Parse(templates["classes.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse classes.gen.go: %w", parseErr))
	}

	protocolsTemplate, parseErr = template.New("protocols.gen.go").Funcs(templateFuncs).Parse(templates["protocols.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse protocols.gen.go: %w", parseErr))
	}

	functionsGenTemplate, parseErr = template.New("functions.gen.go").Funcs(templateFuncs).Parse(templates["functions.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse functions.gen.go: %w", parseErr))
	}

	methodsTemplate, parseErr = template.New("methods.gen.go").Funcs(templateFuncs).Parse(templates["methods.gen.go"])
	if parseErr != nil {
		panic(fmt.Errorf("failed to parse methods.gen.go: %w", parseErr))
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
