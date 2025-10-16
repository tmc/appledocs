// Command method-discovery demonstrates discovering methods from Apple documentation.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"

	"github.com/tmc/appledocs"
)

var (
	cacheDir  = flag.String("cache", filepath.Join(os.Getenv("HOME"), ".appledocs", "cache", "developer.apple.com", "tutorials", "data", "documentation"), "path to cached Apple documentation")
	className = flag.String("class", "", "filter by specific class name")
	framework = flag.String("framework", "", "limit discovery to a specific framework")
	verbose   = flag.Bool("verbose", false, "show detailed method information")
)

func main() {
	flag.Parse()

	// Determine which directory to scan
	scanDir := *cacheDir
	if *framework != "" {
		scanDir = filepath.Join(*cacheDir, *framework)
	}

	// Open the documentation filesystem
	fsys, err := appledocs.Open(scanDir)
	if err != nil {
		log.Fatalf("Failed to open documentation: %v", err)
	}

	// Discover all methods
	fmt.Println("Discovering methods from Apple documentation...")
	classes, err := appledocs.DiscoverMethods(fsys)
	if err != nil {
		log.Fatalf("Failed to discover methods: %v", err)
	}

	fmt.Printf("Discovered %d classes\n\n", len(classes))

	// If a specific class is requested, show only that
	if *className != "" {
		if classMethods, ok := classes[*className]; ok {
			printClassMethods(classMethods, *verbose)
		} else {
			log.Fatalf("Class %s not found", *className)
		}
		return
	}

	// Otherwise, show summary of all classes
	printSummary(classes)
}

func printSummary(classes map[string]*appledocs.ClassMethods) {
	// Sort class names for consistent output
	var classNames []string
	for name := range classes {
		classNames = append(classNames, name)
	}
	sort.Strings(classNames)

	fmt.Println("Class Summary:")
	fmt.Println("=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=")
	fmt.Printf("%-50s %8s %8s %8s %8s\n", "Class", "Instance", "Class", "Props", "Init")
	fmt.Println("-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-" + "-")

	for _, name := range classNames {
		classMethods := classes[name]
		fmt.Printf("%-50s %8d %8d %8d %8d\n",
			truncate(name, 50),
			len(classMethods.InstanceMethods),
			len(classMethods.ClassMethods),
			len(classMethods.Properties),
			len(classMethods.Initializers),
		)
	}

	// Print totals
	var totalIM, totalCM, totalProps, totalInit int
	for _, classMethods := range classes {
		totalIM += len(classMethods.InstanceMethods)
		totalCM += len(classMethods.ClassMethods)
		totalProps += len(classMethods.Properties)
		totalInit += len(classMethods.Initializers)
	}

	fmt.Println("=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=" + "=")
	fmt.Printf("%-50s %8d %8d %8d %8d\n", "TOTAL", totalIM, totalCM, totalProps, totalInit)
}

func printClassMethods(classMethods *appledocs.ClassMethods, verbose bool) {
	fmt.Printf("Class: %s\n", classMethods.ClassName)
	fmt.Printf("External ID: %s\n\n", classMethods.ClassExternalID)

	// Print instance methods
	if len(classMethods.InstanceMethods) > 0 {
		fmt.Printf("Instance Methods (%d):\n", len(classMethods.InstanceMethods))
		fmt.Println("-------------------")
		for _, method := range classMethods.InstanceMethods {
			printMethod(method, verbose)
		}
		fmt.Println()
	}

	// Print class methods
	if len(classMethods.ClassMethods) > 0 {
		fmt.Printf("Class Methods (%d):\n", len(classMethods.ClassMethods))
		fmt.Println("-----------------")
		for _, method := range classMethods.ClassMethods {
			printMethod(method, verbose)
		}
		fmt.Println()
	}

	// Print properties
	if len(classMethods.Properties) > 0 {
		fmt.Printf("Properties (%d):\n", len(classMethods.Properties))
		fmt.Println("-----------")
		for _, prop := range classMethods.Properties {
			printMethod(prop, verbose)
		}
		fmt.Println()
	}

	// Print initializers
	if len(classMethods.Initializers) > 0 {
		fmt.Printf("Initializers (%d):\n", len(classMethods.Initializers))
		fmt.Println("-------------")
		for _, init := range classMethods.Initializers {
			printMethod(init, verbose)
		}
		fmt.Println()
	}
}

func printMethod(method appledocs.MethodInfo, verbose bool) {
	if verbose {
		fmt.Printf("  Name: %s\n", method.Name)
		fmt.Printf("    Title: %s\n", method.Title)
		fmt.Printf("    Kind: %s\n", method.Kind)
		fmt.Printf("    External ID: %s\n", method.ExternalID)
		if method.Signature != "" {
			fmt.Printf("    Signature: %s\n", method.Signature)
		}
		if len(method.Abstract) > 0 {
			fmt.Printf("    Abstract: ")
			for _, content := range method.Abstract {
				fmt.Print(content.Text)
			}
			fmt.Println()
		}
		fmt.Println()
	} else {
		// Compact format
		if method.Signature != "" {
			fmt.Printf("  %s\n", method.Signature)
		} else {
			fmt.Printf("  %s\n", method.Title)
		}
	}
}

func truncate(s string, maxLen int) string {
	if len(s) <= maxLen {
		return s
	}
	return s[:maxLen-3] + "..."
}
