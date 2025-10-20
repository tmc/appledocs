package main

import (
	"fmt"
	"runtime"

	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

func main() {
	// Lock to main thread for Foundation/UTI operations
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	fmt.Println("UniformTypeIdentifiers Examples")
	fmt.Println("================================")

	// Example 1: Create UTTypes from file extensions
	fmt.Println("\n1. Creating UTTypes from File Extensions:")

	pdfType := uniformtypeidentifiers.NewUTTypeWithFilenameExtension("pdf")
	fmt.Printf("   .pdf extension: %v\n", pdfType)

	jpgType := uniformtypeidentifiers.NewUTTypeWithFilenameExtension("jpg")
	fmt.Printf("   .jpg extension: %v\n", jpgType)

	txtType := uniformtypeidentifiers.NewUTTypeWithFilenameExtension("txt")
	fmt.Printf("   .txt extension: %v\n", txtType)

	goType := uniformtypeidentifiers.NewUTTypeWithFilenameExtension("go")
	fmt.Printf("   .go extension: %v\n", goType)

	// Example 2: Create UTTypes from MIME types
	fmt.Println("\n2. Creating UTTypes from MIME Types:")

	jsonMime := uniformtypeidentifiers.NewUTTypeWithMIMEType("application/json")
	fmt.Printf("   application/json: %v\n", jsonMime)

	htmlMime := uniformtypeidentifiers.NewUTTypeWithMIMEType("text/html")
	fmt.Printf("   text/html: %v\n", htmlMime)

	pngMime := uniformtypeidentifiers.NewUTTypeWithMIMEType("image/png")
	fmt.Printf("   image/png: %v\n", pngMime)

	// Example 3: Create UTTypes from identifiers
	fmt.Println("\n3. Creating UTTypes from Type Identifiers:")

	publicText := uniformtypeidentifiers.NewUTTypeWithIdentifier("public.text")
	fmt.Printf("   public.text: %v\n", publicText)

	publicImage := uniformtypeidentifiers.NewUTTypeWithIdentifier("public.image")
	fmt.Printf("   public.image: %v\n", publicImage)

	publicData := uniformtypeidentifiers.NewUTTypeWithIdentifier("public.data")
	fmt.Printf("   public.data: %v\n", publicData)

	// Example 4: Common file types
	fmt.Println("\n4. Common File Type Extensions:")

	extensions := []string{
		"md", "json", "xml", "csv",
		"png", "gif", "svg", "mp4",
		"zip", "tar", "gz",
		"c", "cpp", "swift", "py", "rs",
	}

	for _, ext := range extensions {
		utType := uniformtypeidentifiers.NewUTTypeWithFilenameExtension(ext)
		fmt.Printf("   .%-6s -> %v\n", ext, utType)
	}

	// Example 5: Document types
	fmt.Println("\n5. Document Types:")

	docTypes := map[string]string{
		"docx": "Microsoft Word",
		"xlsx": "Microsoft Excel",
		"pptx": "Microsoft PowerPoint",
		"pages": "Apple Pages",
		"numbers": "Apple Numbers",
		"keynote": "Apple Keynote",
	}

	for ext, desc := range docTypes {
		utType := uniformtypeidentifiers.NewUTTypeWithFilenameExtension(ext)
		fmt.Printf("   .%-10s (%-25s): %v\n", ext, desc, utType)
	}

	// Example 6: Media types
	fmt.Println("\n6. Media File Types:")

	mediaTypes := []struct {
		ext  string
		desc string
	}{
		{"mp3", "Audio"},
		{"m4a", "Audio"},
		{"wav", "Audio"},
		{"mp4", "Video"},
		{"mov", "Video"},
		{"avi", "Video"},
	}

	for _, media := range mediaTypes {
		utType := uniformtypeidentifiers.NewUTTypeWithFilenameExtension(media.ext)
		fmt.Printf("   .%-6s (%-6s): %v\n", media.ext, media.desc, utType)
	}

	fmt.Println("\n✓ All UniformTypeIdentifiers operations completed successfully!")
	fmt.Println("\nNote: UTType objects are used throughout macOS/iOS for:")
	fmt.Println("  - File type identification")
	fmt.Println("  - Drag & drop operations")
	fmt.Println("  - File open/save dialogs")
	fmt.Println("  - Document type declarations")

	// Suppress unused variable warnings
	_ = pdfType
	_ = jpgType
	_ = txtType
	_ = goType
	_ = jsonMime
	_ = htmlMime
	_ = pngMime
	_ = publicText
	_ = publicImage
	_ = publicData
}
