// Package main demonstrates cross-referencing between Swift and Objective-C APIs.
package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/tmc/appledocs"
)

func main() {
	// Open the documentation cache
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	cacheDir := filepath.Join(homeDir, ".appledocs/cache/developer.apple.com/tutorials/data/documentation")

	fsys, err := appledocs.Open(cacheDir)
	if err != nil {
		log.Fatal(err)
	}

	// Example 1: Get cross-reference for NSString
	fmt.Println("=== NSString Cross-Reference ===")
	doc, err := appledocs.GetSymbol(fsys, "Foundation/NSString")
	if err != nil {
		log.Fatal(err)
	}

	ref := appledocs.GetCrossReference(doc)
	fmt.Printf("Symbol: %s\n", ref.Title)
	fmt.Printf("Kind: %s\n\n", ref.SymbolKind)

	if ref.Available.Swift {
		fmt.Println("Swift Declaration:")
		fmt.Printf("  %s\n\n", ref.SwiftDeclaration)
	}

	if ref.Available.ObjectiveC {
		fmt.Println("Objective-C Declaration:")
		fmt.Printf("  %s\n\n", ref.ObjCDeclaration)
	}

	// Example 2: Check which language variants are available
	fmt.Println("=== Language Variant Availability ===")
	symbols := []string{
		"Foundation/NSArray",
		"Foundation/NSDate",
		"Foundation/NSURL",
	}

	for _, symbol := range symbols {
		doc, err := appledocs.GetSymbol(fsys, symbol)
		if err != nil {
			fmt.Printf("%-30s: Error - %v\n", symbol, err)
			continue
		}

		hasSwift := appledocs.HasSwiftVariant(doc)
		hasObjC := appledocs.HasObjectiveCVariant(doc)

		fmt.Printf("%-30s: Swift=%v ObjC=%v\n", symbol, hasSwift, hasObjC)
	}

	// Example 3: Get specific language variants
	fmt.Println("\n=== Getting Specific Language Variants ===")
	doc, err = appledocs.GetSymbol(fsys, "Foundation/NSArray")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Original language: %s\n\n", doc.Identifier.InterfaceLanguage)

	// Get Swift variant
	swiftDoc, err := appledocs.GetSwiftVariant(doc)
	if err != nil {
		fmt.Printf("Swift variant error: %v\n", err)
	} else {
		fmt.Println("Swift variant obtained successfully")
		fmt.Printf("Language: %s\n", swiftDoc.Identifier.InterfaceLanguage)
		fmt.Printf("Declaration: %s\n\n", appledocs.GetDeclarationText(swiftDoc, appledocs.LanguageSwift))
	}

	// Get Objective-C variant
	objcDoc, err := appledocs.GetObjectiveCVariant(doc)
	if err != nil {
		fmt.Printf("Objective-C variant error: %v\n", err)
	} else {
		fmt.Println("Objective-C variant obtained successfully")
		fmt.Printf("Language: %s\n", objcDoc.Identifier.InterfaceLanguage)
		fmt.Printf("Declaration: %s\n", appledocs.GetDeclarationText(objcDoc, appledocs.LanguageObjectiveC))
	}
}
