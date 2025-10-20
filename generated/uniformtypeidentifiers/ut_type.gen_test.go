// Code generated from Apple documentation for UniformTypeIdentifiers. DO NOT EDIT.

package uniformtypeidentifiers_test

import (
	"github.com/tmc/appledocs/generated/uniformtypeidentifiers"
)

// Suppress unused import errors
var _ = uniformtypeidentifiers.NewUTType


// ExampleNewUTTypeExportedTypeWithIdentifier demonstrates how to create a UTType instance using NewUTTypeExportedTypeWithIdentifier.
// Creates a type your app owns based on an identifier.
func ExampleNewUTTypeExportedTypeWithIdentifier() {
	_ = uniformtypeidentifiers.NewUTTypeExportedTypeWithIdentifier(
		"identifier", // identifier string
	)
	// Output:
}



// ExampleNewUTTypeWithMIMEType demonstrates how to create a UTType instance using NewUTTypeWithMIMEType.
// Creates a type based on a MIME type.
func ExampleNewUTTypeWithMIMEType() {
	_ = uniformtypeidentifiers.NewUTTypeWithMIMEType(
		"mimeType", // mimeType string
	)
	// Output:
}


// ExampleNewUTTypeWithIdentifier demonstrates how to create a UTType instance using NewUTTypeWithIdentifier.
// Creates a type based on an identifier.
func ExampleNewUTTypeWithIdentifier() {
	_ = uniformtypeidentifiers.NewUTTypeWithIdentifier(
		"identifier", // identifier string
	)
	// Output:
}

// ExampleNewUTTypeWithFilenameExtension demonstrates how to create a UTType instance using NewUTTypeWithFilenameExtension.
// Creates a type that represents the specified filename extension.
func ExampleNewUTTypeWithFilenameExtension() {
	_ = uniformtypeidentifiers.NewUTTypeWithFilenameExtension(
		"filenameExtension", // filenameExtension string
	)
	// Output:
}


// ExampleNewUTTypeImportedTypeWithIdentifier demonstrates how to create a UTType instance using NewUTTypeImportedTypeWithIdentifier.
// Creates a type your app uses, but doesn’t own, based on an identifier.
func ExampleNewUTTypeImportedTypeWithIdentifier() {
	_ = uniformtypeidentifiers.NewUTTypeImportedTypeWithIdentifier(
		"identifier", // identifier string
	)
	// Output:
}


